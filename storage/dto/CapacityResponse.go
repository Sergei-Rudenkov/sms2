package dto

type CapacityResponse struct{
	Val int
	err error
	TransactionID string
}

func (r *CapacityResponse) GetError() (error) {
	return r.err
}

func (r *CapacityResponse) SetError(err error)  {
	r.err = err
}

func (r *CapacityResponse) Read() interface{}  {
	return r.Val
}