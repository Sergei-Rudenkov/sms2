package server

import (
	"net/http"
	"sms2/storage"
	"errors"
	"fmt"
	"sms2/storage/dto"
	_ "github.com/google/uuid"
	"math/rand"
	"strconv"
)

var responseChan chan dto.Responder

func Start() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	cr := request2CacheRequest(r)
	cr.TransactionID = string(rand.Int())
	responder := processReq(cr)
	switch responder.(type) {
	case *dto.KeysResponse:
		fmt.Fprintf(w, responder.Read().(string))
	case *dto.GetResponse:
		fmt.Fprintf(w, responder.Read().(string))
	case *dto.RemoveResponse:
		fmt.Fprintf(w, fmt.Sprintf("{'succ': %s}", strconv.FormatBool(responder.(*dto.RemoveResponse).Success)))
	case *dto.SetResponse:
		fmt.Fprintf(w, fmt.Sprintf("{'succ': %s}", strconv.FormatBool(responder.(*dto.SetResponse).Evicted)))
	case *dto.CapacityResponse:
		if i, ok := responder.Read().(int); ok {
			fmt.Fprintf(w, fmt.Sprintf("{'capasity': %d, 'err': %s}", i, responder.GetError()))
		}
	case *dto.EmptyResponse:
		fmt.Fprintf(w, fmt.Sprintf("{'err': '%s'}", responder.GetError().Error()))
	}
}

func processReq(cr *dto.CacheRequest) (r dto.Responder) {
	switch cr.Operation {
	case "set":
		r = storage.Set(cr)
	case "get":
		r = storage.Get(cr)
	case "remove":
		r = storage.Remove(cr)
	case "keys":
		r = storage.Keys()
	case "capacity":
		r = storage.Capacity()
	default:
		r = &dto.EmptyResponse{
			Err:           errors.New("Unknown opperation"),
			TransactionID: cr.TransactionID,
		}
	}
	return r
}

func request2CacheRequest(r *http.Request) *dto.CacheRequest {
	if err := r.ParseForm(); err != nil {
		fmt.Println("ParseForm() err: %v", err)
		return nil
	}
	return &dto.CacheRequest{
		r.FormValue("operation"),
		r.FormValue("key"),
		r.FormValue("value"),
		r.FormValue("valueType"),
		r.FormValue("transactionID"),
	}
}
