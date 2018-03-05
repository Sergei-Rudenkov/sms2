package server

import (
	"net/http"
)

func ServeHttp() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

}

func processReq() () {
	switch `` {
	case "set":
	case "get":
	case "remove":
	case "keys":
	case "capacity":
	default:
	}
}

func request2CacheRequest(r *http.Request) {
	r.FormValue("operation")
	r.FormValue("key")
	r.FormValue("value")
}
