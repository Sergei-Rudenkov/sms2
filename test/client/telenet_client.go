package client

import (
	"github.com/reiver/go-telnet"
	log "github.com/inconshreveable/log15"
	"fmt"
)


func Set(key, value string, ttl int) {
	conn, err := telnet.DialTo("localhost:5555")
	if err != nil {
		log.Error(err.Error())
	}
	_, err = conn.Write([]byte(fmt.Sprintf("set %s %s %d", key, value, ttl)))
	if err != nil {
		log.Error(err.Error())
	}
	_, err = conn.Write([]byte("\n"))
	if err != nil {
		log.Error(err.Error())
	}
}

// TODO implement:
// Get, Keys, Capacity,
// LSet, LAdd, LGet
// DSet, DAdd, DRemove

