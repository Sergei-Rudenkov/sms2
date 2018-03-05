package storage

import (
	"time"
	"sync"
	"sms2/storage/provider/fixed"
	"strconv"
	"sms2/storage/provider/agile"
	"sms2/storage/provider"
)

var (
	singleton provider.Cache
	once      sync.Once
	cacheType string
)

func GetCacheProviderType() string{
	return cacheType
}

func GetCache() provider.Cache {
	return singleton
}

func InitCache(t string, arg ...string)  {
	once.Do(func() {
		switch t {
		case `fixed`:
			capacity, err := strconv.Atoi(arg[0])
			if err != nil {
				panic("Can not parse int from first argument which is should be `Capacity`")
			}
			tis, err := strconv.Atoi(arg[1])
			if err != nil {
				panic("Can not parse int from second argument which is should be `Time in seconds`")
			}
			singleton = fixed.New(capacity, fixed.WithTTL(time.Duration(tis)*time.Second))
		case `agile`:
			singleton = agile.New()
		}
		cacheType = t
	})
}

