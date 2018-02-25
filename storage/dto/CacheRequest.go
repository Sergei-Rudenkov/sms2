package dto

type CacheRequest struct {
	Operation,
	Key,
	Value,
	ValueType,
	TransactionID string
}
