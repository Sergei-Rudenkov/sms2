package server

import (
	"net/http"
	"strconv"
	"fmt"
	"sms2/storage"
	"sms2/util"
	"sms2/service"
)

func ServeHttp(port string) {
	http.HandleFunc("/", handler)
	http.ListenAndServe(port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	operation, key, value, ttl, err := service.SHttpRequestParamParser(r)
	switch operation {
	case "set":
		if err != nil{
			fmt.Fprintf(w, "%s", err.Error())
			return
		}
		evicted := storage.GetCache().Set(key, value,  ttl)
		fmt.Fprintf(w, strconv.FormatBool(evicted))
	case "get":
		value, exist := storage.GetCache().Get(key)
		if !exist {
			fmt.Fprintf(w, "%s", "value for this key does not exist")
			return
		}
		fmt.Fprintf(w, "%s", value)
	case "remove":
		ok := storage.GetCache().Del(key)
		fmt.Fprintf(w, "%s", strconv.FormatBool(ok))
	case "keys":
		listOfKeys := storage.GetCache().Keys()
		keys := util.ListOfObjectsToConcatString(listOfKeys)
		fmt.Fprintf(w, "%s", keys)
	case "capacity":
		capacity := storage.GetCache().Cap()
		fmt.Fprintf(w, "%d", capacity)
	default:
		fmt.Fprintf(w, "%s", "unknown operation")
	}
}
