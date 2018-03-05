package server

import (
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
	"io"
	log "github.com/inconshreveable/log15"
	"sms2/storage"
	"errors"
	"strconv"
	"fmt"
	"strings"
	"time"
)

func ServeTelnetConnection() {
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

	// Register the "capacity" command.
	commandName     = "get"
	commandProducer = telsh.ProducerFunc(getProducer)
	shellHandler.Register(commandName, commandProducer)

	// Register the "remove" command.
	commandName     = "remove"
	commandProducer = telsh.ProducerFunc(removeProducer)
	shellHandler.Register(commandName, commandProducer)

	addr := ":5555"
	if err := telnet.ListenAndServe(addr, shellHandler); nil != err {
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
		oi.LongWriteString(stdout, 	"{" + strings.Join(stringList,",") + "}")
		return nil
	})
}

func setProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	log.Info("`set` command received","args:", args)
	argMap, err := argumentParser(name, args...)
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
	argMap, err := argumentParser(name, args...)
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
	argMap, err := argumentParser(name, args...)
	ok := storage.GetCache().Del(argMap[`key`])

	return telsh.PromoteHandlerFunc(func(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
		if err != nil{
			log.Debug(err.Error())
			oi.LongWriteString(stderr, err.Error())
			return nil
		}
		if !ok {
			oi.LongWriteString(stdout,"Value for this key does not exist. Nothing was removed.")
			return nil
		}
		oi.LongWriteString(stdout, "Removed.")
		return nil
	})
}

func argumentParser (commandName string, args ...string) (map[string]string, error) {
	var err error
	argMap := make(map[string]string)
	cacheProvider := storage.GetCacheProviderType()
	switch commandName {
	case `set`:
		if len(args) >= 2 {
			argMap[`key`] = args[0]
			argMap[`value`] = args[1]
		} else {
			err = errors.New("absence of key and(or) value argument(s) in `set` operation")
		}
		if cacheProvider == `agile` && len(args) >= 3{
			if intTtl, intConversionErr := strconv.Atoi(args[2]); intConversionErr == nil && intTtl > 0 {
				argMap[`ttl`] = args[2]
				return argMap, err
			}
			err = errors.New("please, provide int ttl greater then 0")
		} else {
			err = errors.New("your cache provider is `Agile`, ttl is obligatory as the 3rd argument")
		}
	case `get`, `remove`:
		if args[0] != ``{
			argMap[`key`] = args[0]
			return argMap, err
		}
		err = errors.New("absence of key argument in `get` or `remove` operation")
	}
	return argMap, err
}