package dto

import "fmt"

type GetResponse struct {
	Value         interface{}
	err           error
	Success		  bool
	TransactionID string
}

func (r *GetResponse) GetError() (error) {
	return r.err
}

func (r *GetResponse) SetError(err error) {
	r.err = err
}

func (r *GetResponse) Read() interface{} {
	return fmt.Sprintf("{'value': '%s', 'type': '%s', 'err', '%s'}", r.Value, r.err.Error())
}
