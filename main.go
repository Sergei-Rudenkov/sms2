package main

import (
	"sms2/server"
	log "github.com/inconshreveable/log15"
)


func main ()  {
	customiseLogging()
	log.Info("SMS2 get starting...")

	go func(){
		log.Info("Start listening on http")
		server.ServeHttp()
	}()

	log.Info("Start listening telnet")
	server.ServeTelnetConnection()

}


func customiseLogging() {
	h := log.CallerFileHandler(log.StdoutHandler)
	log.Root().SetHandler(h)
}
