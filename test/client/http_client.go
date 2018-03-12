package client

import (
	"net/http"
	"net/url"
	log "github.com/inconshreveable/log15"
	"strconv"
	"io/ioutil"
	"fmt"
)

type HttpClient struct {
	Host string
	Port int
}

type Connector interface{
	HttpSet(key, value string, ttl int)
	HttpGet(key string) string
	HttpRemove(key string) bool
}


// HttpSet - ttl param is required only for agile cache type
func (c *HttpClient) HttpSet(key, value string, ttl int) {
	_, err := http.PostForm(fmt.Sprintf("http://%s:%d", c.Host, c.Port),
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

func (c *HttpClient) HttpGet(key string) string  {
	resp, err := http.PostForm(fmt.Sprintf("http://%s:%d", c.Host, c.Port),
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

func (c *HttpClient)HttpRemove(key string) bool{
	resp, err := http.PostForm(fmt.Sprintf("http://%s:%d", c.Host, c.Port),
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

// TODO implement:
// LSet, LAdd, LGet
// DSet, DAdd, DRemove

