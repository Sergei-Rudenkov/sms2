package dto

type Responder interface {
	Read() (interface{})
	GetError() (error)
	SetError(error)
	GetTransactionId() (string)
}
