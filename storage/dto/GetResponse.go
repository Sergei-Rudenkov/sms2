package dto

import "fmt"

type GetResponse struct {
	Value         interface{}
	err           error
	Success       bool
	TransactionID string
}

func (r *GetResponse) GetError() (error) {
	return r.err
}

func (r *GetResponse) SetError(err error) {
	r.err = err
}

func (r *GetResponse) Read() interface{} {
	if (r.Value != nil && r.err != nil) {
		return fmt.Sprintf("{'value': '%s', 'err', '%s'}", r.Value.(string), r.err.Error())
	} else if(r.err == nil && r.Value != nil){
		return fmt.Sprintf("{'value': '%s'}", r.Value.(string))
	} else if(r.Value == nil && r.err != nil){
		return fmt.Sprintf("{'err': '%s'}", r.err.Error())
	}
	return "{'value': nil}"
}

