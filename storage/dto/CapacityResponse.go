package dto

type CapacityResponse struct{
	Val int
	Err error
	TransactionID string
}

func (r *CapacityResponse) GetError() (error) {
	return r.Err
}

func (r *CapacityResponse) SetError(err error)  {
	r.Err = err
}

func (r *CapacityResponse) Read() interface{}  {
	return r.Val
}

func (r *CapacityResponse) GetTransactionId() string  {
	return r.TransactionID
}