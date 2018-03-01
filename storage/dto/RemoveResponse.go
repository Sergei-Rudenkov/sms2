package dto

type RemoveResponse struct{
	err error
	Success bool
	TransactionID string
}

func (r *RemoveResponse) GetError() (error) {
	return r.err
}

func (r *RemoveResponse) SetError(err error) {
	r.err = err
}

func (r *RemoveResponse) Read() interface{} {
	return ""
}