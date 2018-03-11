package client

import (
	"net/http"
	"net/url"
	log "github.com/inconshreveable/log15"
	"strconv"
	"io/ioutil"
)

//////////////////////////////////////////////////////////////
//
// Http client supports only operations with string values,
// for list and dictionary support use telnet client
//
//////////////////////////////////////////////////////////////

func HttpSet(key, value string, ttl int) {
	_, err := http.PostForm("localhost:8080",
		url.Values{
			"operation": {"set"},
			"key": {key},
			"value": {value},
			"ttl": {strconv.Itoa(ttl)},
			})
	if err != nil{
		log.Error(err.Error())
	}
}

func HttpFixedSet(key, value string) {
	_, err := http.PostForm("localhost:8080",
		url.Values{
			"operation": {"set"},
			"key": {key},
			"value": {value},
		})
	if err != nil{
		log.Error(err.Error())
	}
}

func HttpGet(key string) string  {
	resp, err := http.PostForm("localhost:8080",
		url.Values{
			"operation": {"get"},
			"key": {key},
		})
	if err != nil{
		log.Error(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func HttpRemove(key string) bool{
	resp, err := http.PostForm("localhost:8080",
		url.Values{
			"operation": {"remove"},
			"key": {"f"},
		})
	if err != nil{
		log.Error(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Error(err.Error())
	}
	succ, err := strconv.ParseBool(string(body))
	if err != nil{
		log.Error(err.Error())
	}
	return succ
}

