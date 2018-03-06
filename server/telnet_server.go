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
	"strings"
	"time"
)

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

	if err := telnet.ListenAndServe(port, shellHandler); nil != err {
		panic(err)
	}
}


func keysProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	log.Info("`keys` command received")
	keys := storage.GetCache().Keys()
	stringList := make([]string, len(keys))
	for i := range keys {
		stringList[i] = keys[i].(string)
	}

	return telsh.PromoteHandlerFunc(func(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
		oi.LongWriteString(stdout, 	strings.Join(stringList,","))
		return nil
	})
}

func setProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	log.Info("`set` command received","args:", args)
	argMap, err := util.TelnetArgumentParser(name, args...)
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
	argMap, err := util.TelnetArgumentParser(name, args...)
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
		oi.LongWriteString(stdout, value.(string))
		return nil
	})
}

func removeProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	log.Info("`remove` command received","args:", args)
	argMap, err := util.TelnetArgumentParser(name, args...)
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