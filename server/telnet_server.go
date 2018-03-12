package server

import (
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
	"io"
	log "github.com/inconshreveable/log15"
	"sms2/storage"
	"sms2/util"
	"strconv"
	"fmt"
	"time"
	"sms2/service"
	"strings"
	"errors"
)

// ServeTelnetConnection - start listening to telnet connections.
// register commandProducers
func ServeTelnetConnection(port string) {
	shellHandler := telsh.NewShellHandler()
	shellHandler.WelcomeMessage = `
/_____/\ /__//_//_/\ /_____/\ /_____/\     
\::::_\/_\::\| \| \ \\::::_\/_\:::_:\ \    
 \:\/___/\\:.      \ \\:\/___/\   _\:\|    
  \_::._\:\\:.\-/\  \ \\_::._\:\ /::_/__   
    /____\:\\. \  \  \ \ /____\:\\:\____/\ 
    \_____\/ \__\/ \__\/ \_____\/ \_____\/ 
`
	// Register the "set" command.
	commandName     := "set"
	commandProducer := telsh.ProducerFunc(setProducer)
	shellHandler.Register(commandName, commandProducer)

	// Register the "keys" command.
	commandName     = "keys"
	commandProducer = telsh.ProducerFunc(keysProducer)
	shellHandler.Register(commandName, commandProducer)

	// Register the "capacity" command.
	commandName     = "capacity"
	commandProducer = telsh.ProducerFunc(capacityProducer)
	shellHandler.Register(commandName, commandProducer)

	// Register the "get" command.
	commandName     = "get"
	commandProducer = telsh.ProducerFunc(getProducer)
	shellHandler.Register(commandName, commandProducer)

	// Register the "remove" command.
	commandName     = "remove"
	commandProducer = telsh.ProducerFunc(removeProducer)
	shellHandler.Register(commandName, commandProducer)

	////////////////////////////////
	//
	// list support
	//
	////////////////////////////////

	// Register the "lset" command.
	commandName     = "lset"
	commandProducer = telsh.ProducerFunc(lsetProducer)
	shellHandler.Register(commandName, commandProducer)

	// Register the "lget" command.
	commandName     = "lget"
	commandProducer = telsh.ProducerFunc(lgetProducer)
	shellHandler.Register(commandName, commandProducer)

	// Register the "ladd" command.
	commandName     = "ladd"
	commandProducer = telsh.ProducerFunc(laddProducer)
	shellHandler.Register(commandName, commandProducer)

	if err := telnet.ListenAndServe(port, shellHandler); nil != err {
		panic(err)
	}
}


func keysProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	log.Info("`keys` command received")
	listOfKeys := storage.GetCache().Keys()
	keys := util.ListOfObjectsToConcatString(listOfKeys)

	return telsh.PromoteHandlerFunc(func(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
		oi.LongWriteString(stdout, 	keys)
		return nil
	})
}

func setProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	log.Info("`set` command received","args:", args)
	argMap, err := service.STelnetArgumentParser(name, args...)
	tis, _ := strconv.Atoi(argMap[`ttl`]) // error should already have been checked in argumentParser function
	evicted := storage.GetCache().Set(argMap[`key`], argMap[`value`],  time.Duration(tis)* time.Second)

	return telsh.PromoteHandlerFunc(func(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
		if err != nil{
			log.Debug(err.Error())
			oi.LongWriteString(stderr, err.Error())
			return nil
		}
		oi.LongWriteString(stdout, fmt.Sprintf("Evicted: %s", strconv.FormatBool(evicted)))
		return nil
	})
}

func capacityProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	log.Info("`capacity` command received","args:", args)
	capacity := storage.GetCache().Cap()

	return telsh.PromoteHandlerFunc(func(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
		oi.LongWriteString(stdout, strconv.Itoa(capacity))
		return nil
	})
}

func getProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	log.Info("`get` command received","args:", args)
	argMap, err := service.STelnetArgumentParser(name, args...)
	value, exist := storage.GetCache().Get(argMap[`key`])

	return telsh.PromoteHandlerFunc(func(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
		if err != nil{
			log.Debug(err.Error())
			oi.LongWriteString(stderr, err.Error())
			return nil
		}
		if !exist {
			oi.LongWriteString(stdout,"Value for this key does not exist.")
			return nil
		}
		switch value.(type) {
		case string:
			oi.LongWriteString(stdout, value.(string))
		case []string:
			oi.LongWriteString(stdout, strings.Join(value.([]string), `,`))
		}
		return nil
	})
}

func removeProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	log.Info("`remove` command received","args:", args)
	argMap, err := service.STelnetArgumentParser(name, args...)
	ok := storage.GetCache().Del(argMap[`key`])

	return telsh.PromoteHandlerFunc(func(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
		if err != nil{
			log.Debug(err.Error())
			oi.LongWriteString(stderr, err.Error())
			return nil
		}
		if !ok {
			log.Debug("Value for this key does not exist. Nothing was removed.")
		}
		oi.LongWriteString(stdout, strconv.FormatBool(ok))
		return nil
	})
}

////////////////////////////////
//
// list support
//
////////////////////////////////

func lsetProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	log.Info("`lset` command received","args:", args)
	argMap, err := service.STelnetArgumentParser(name, args...)
	tis, _ := strconv.Atoi(argMap[`ttl`]) // error should already have been checked in argumentParser function
	listValue := util.StringToList(argMap[`value`]) // make list from string like [1,2,3]
	evicted := storage.GetCache().Set(argMap[`key`], listValue,  time.Duration(tis)* time.Second)

	return telsh.PromoteHandlerFunc(func(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
		if err != nil{
			log.Debug(err.Error())
			oi.LongWriteString(stderr, err.Error())
			return nil
		}
		oi.LongWriteString(stdout, fmt.Sprintf("Evicted: %s", strconv.FormatBool(evicted)))
		return nil
	})
}

func lgetProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	log.Info("`lget` command received","args:", args)
	argMap, err := service.LTelnetArgumentParser(name, args...)
	if err != nil{
		log.Error("error during lget argument parsing ", err.Error())
	}
	value, exist := storage.GetCache().Get(argMap[`key`])

	return telsh.PromoteHandlerFunc(func(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
		var lreturn []string
		if exist {
			list := value.([]string)
			if argMap[`first`] != `` && argMap[`last`] != ``{
				first, _ := strconv.Atoi(argMap[`first`]) // error check skipped bc regexp already proved int
				last, _ := strconv.Atoi(argMap[`last`]) // error check skipped bc regexp already proved int
				lreturn = list[first:last]
			}
			if argMap[`first`] == `` && argMap[`last`] != `` {
				last, _ := strconv.Atoi(argMap[`last`]) // error check skipped bc regexp already proved int
				lreturn = list[:last]
			}
			if argMap[`last`] == `` && argMap[`first`] != `` {
				first, _ := strconv.Atoi(argMap[`first`]) // error check skipped bc regexp already proved int
				lreturn = list[first:]
			}
			oi.LongWriteString(stdout, strings.Join(lreturn, ","))
			return nil
		}
		if err != nil{
			log.Error("error during lget argument parsing ", err.Error())
			oi.LongWriteString(stderr, err.Error())
			return nil
		}
		oi.LongWriteString(stdout, strconv.FormatBool(exist))
		return nil
	})
}

func laddProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	var evicted bool
	log.Info("`ladd` command received","args:", args)
	argMap, err := service.LTelnetArgumentParser(name, args...)
	value, exist := storage.GetCache().Get(argMap[`key`])
	tis, _ := strconv.Atoi(argMap[`ttl`]) // error should already have been checked in argumentParser function
	if exist {
		switch value.(type){
		case []string:
			list := value.([]string)
			list = append(list, argMap[`item`])
			evicted = storage.GetCache().Set(argMap[`key`], list, time.Duration(tis)* time.Second)
		default:
			err = errors.New("value for passed key is not a list. Cannot accomplish ladd")
		}
	}

	return telsh.PromoteHandlerFunc(func(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
		if err != nil{
			log.Debug(err.Error())
			oi.LongWriteString(stderr, err.Error())
			return nil
		}
		oi.LongWriteString(stdout, fmt.Sprintf("Evicted: %s", strconv.FormatBool(evicted)))
		return nil
	})

	////////////////////////////////
	//
	// dictionary support
	// (is not implemented yet)
	//
	////////////////////////////////
}