package main

import (
	"sms2/server"
	log "github.com/inconshreveable/log15"
	"sms2/storage"
	"os"
)


func main ()  {
	initCache(os.Args)
	customiseLogging()

	log.Info("SMS2 get starting...")
	go func(){
		log.Info("Start listening on http")
		server.ServeHttp()
	}()

	log.Info("Start listening telnet")
	server.ServeTelnetConnection()
}

func initCache(args []string) {
	var cacheType, capacity, ttl string
	if len(args) >= 2{
		cacheType = args[1]
	}else {
		panic("please, specify cache type as first program argument")
	}
	if len(args) >= 4{
		capacity = args[2] // only for type `fixed`
		ttl = args[3] // only for type `fixed`
	}
	storage.InitCache(cacheType, capacity, ttl)
}


func customiseLogging() {
	h := log.CallerFileHandler(log.StdoutHandler)
	log.Root().SetHandler(h)
}
