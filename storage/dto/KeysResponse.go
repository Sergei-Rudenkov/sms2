package dto

type KeysResponse struct{
	List interface{}
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
	return r.List
}