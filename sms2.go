package main

import (
	"sms2/server"
	log "github.com/inconshreveable/log15"
	"sms2/storage"
	"os"
	"encoding/json"
)

type Configuration struct {
	HttpPort   string
	TelnetPort string
}

func main() {
	//setup logging
	h := log.CallerFileHandler(log.StdoutHandler)
	log.Root().SetHandler(h)

	// reading configuration
	file, err := os.Open("config/config.json")
	defer file.Close()
	if err != nil {
		log.Error(err.Error())
	}

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Error(err.Error())
	}

	log.Info("SMS2 initializing cache instance...")
	initCache(os.Args)

	go func() {
		log.Info("Start listening on http")
		server.ServeHttp(configuration.HttpPort)
	}()

	log.Info("Start listening telnet")
	server.ServeTelnetConnection(configuration.TelnetPort)
}

func initCache(args []string) {
	var cacheType, capacity, ttl string
	if len(args) >= 2 {
		cacheType = args[1]
	} else {
		panic("please, specify cache type as first program argument")
	}
	if len(args) >= 4 {
		capacity = args[2] // only for type `fixed`
		ttl = args[3]      // only for type `fixed`
	}
	storage.InitCache(cacheType, capacity, ttl)
}
