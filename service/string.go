package service

import (
	"time"
	"strconv"
	"net/http"
	log "github.com/inconshreveable/log15"
	"errors"
	"sms2/storage"
)

func SHttpRequestParamParser(r *http.Request) (operation, key, value string, ttl time.Duration, err error)  {
	operation = r.FormValue("operation")
	key = r.FormValue("key")
	value = r.FormValue("value")
	ttlString := r.FormValue("ttl")
	if ttlString != `` && storage.GetCacheProviderType() == `agile` {
		tis, err := strconv.Atoi(ttlString)
		if err != nil || tis < 0 {
			log.Error("error to parse int from ttl argument", "err", err.Error())
			err = errors.New("error to parse int from ttl argument or ttl is less then zero")
			return operation, key, value, ttl, err
		}
		ttl = time.Duration(tis)* time.Second
	}else {
		err = errors.New("ttl should be provided as argument for set operation")
	}
	return operation, key, value, ttl, err
}

func STelnetArgumentParser(commandName string, args ...string) (map[string]string, error) {
	var err error
	argMap := make(map[string]string)
	cacheProvider := storage.GetCacheProviderType()
	switch commandName {
	case `set`, `lset`:
		if len(args) >= 2 {
			argMap[`key`] = args[0]
			argMap[`value`] = args[1]
		} else {
			err = errors.New("absence of key and(or) value argument(s) in `set` operation")
		}
		if cacheProvider == `agile` {
			if len(args) >= 3 {
				if intTtl, intConversionErr := strconv.Atoi(args[2]); intConversionErr == nil && intTtl > 0 {
					argMap[`ttl`] = args[2]
					return argMap, err
				}
			}
			err = errors.New("please, provide int ttl greater then 0 as third argument")
		}
	case `get`, `remove`:
		if len(args) > 0 && args[0] != ``{
			argMap[`key`] = args[0]
			return argMap, err
		}
		err = errors.New("absence of key argument in `get` or `remove` operation")
	}
	return argMap, err
}
