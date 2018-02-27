package server

import (
	"net/http"
	"sms2/storage"
	"errors"
	"fmt"
	"sms2/storage/dto"
	_ "github.com/google/uuid"
	"math/rand"
)

var responseChan chan dto.Responder

func Start() {
	responseChan = make(chan dto.Responder)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	cr := request2CacheRequest(r)
	cr.TransactionID = string(rand.Int())
	go func() {
		responder := processReq(cr)
		responseChan <- responder
	}()
	fmt.Fprintf(w, "Hello!!!")

	go func() {
		for r := range responseChan {
			if r.GetTransactionId() == cr.TransactionID {
				switch r.(type) {
				//case dto.KeysResponse:
				//case dto.GetResponse:
				//case dto.RemoveResponse:
				//case dto.SetResponse:
				case *dto.CapacityResponse:
					if i, ok := r.Read().(int); ok {
						fmt.Fprintf(w, fmt.Sprintf("{'capasity': %d, 'err': %s}", i, r.GetError()))
					}
				case *dto.EmptyResponse:
					fmt.Fprintf(w, fmt.Sprintf("{'err': '%s'}"), r.GetError().Error())
				}
			}
		}
	}()
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
			Err: errors.New("Unknown opperation"),
			TransactionID: cr.TransactionID,
		}
	}
	return r
}

func request2CacheRequest(r *http.Request) *dto.CacheRequest {
	return &dto.CacheRequest{
		r.FormValue("operation"),
		r.FormValue("key"),
		r.FormValue("value"),
		r.FormValue("valueType"),
		r.FormValue("transactionID"),
	}
}
