package dto

type SetResponse struct {
	Success       bool
	err           error
	TransactionID string
}

func (r *SetResponse) GetError() (error) {
	return r.err
}

func (r *SetResponse) SetError(err error) {
	r.err = err
}

func (r *SetResponse) Read() interface{} {
	return ""
}
