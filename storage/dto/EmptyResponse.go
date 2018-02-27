package dto

type EmptyResponse struct {
	Err error
	TransactionID string
}

func (r *EmptyResponse) GetError() (error) {
	return r.Err
}

func (r *EmptyResponse) SetError(err error)  {
	r.Err = err
}

func (r *EmptyResponse) Read() interface{}  {
	return r.Err
}

func (r *EmptyResponse) GetTransactionId() string  {
	return r.TransactionID
}
