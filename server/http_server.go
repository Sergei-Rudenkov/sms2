package server

import (
	"net/http"
	"strconv"
	"fmt"
	"sms2/storage"
	"strings"
	"sms2/util"
)

func ServeHttp(port string) {
	http.HandleFunc("/", handler)
	http.ListenAndServe(port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	operation, key, value, ttl, err := util.HttpRequestParamParser(r)
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
		keys := storage.GetCache().Keys()
		stringList := make([]string, len(keys))
		for i := range keys {
			stringList[i] = keys[i].(string)
		}
		fmt.Fprintf(w, "%s", strings.Join(stringList,","))
	case "capacity":
		capacity := storage.GetCache().Cap()
		fmt.Fprintf(w, "%d", capacity)
	default:
		fmt.Fprintf(w, "%s", "unknown operation")
	}
}
