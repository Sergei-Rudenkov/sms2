package dto

import "strings"

type KeysResponse struct{
	List []interface{}
	Success bool
	err error
	TransactionID string
}

func (r *KeysResponse) GetError() (error) {
	return r.err
}

func (r *KeysResponse) SetError(err error) {
	r.err = err
}

func (r *KeysResponse) Read() interface{} {
	stringList := make([]string, len(r.List))
	for i := range r.List {
		stringList[i] = r.List[i].(string)
	}
	return "{" + strings.Join(stringList,",") + "}"
}