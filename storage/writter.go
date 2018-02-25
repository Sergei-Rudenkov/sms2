package storage

import (
	"sms2/storage/dto"
)

func Set (cr *dto.CacheRequest) (r dto.Responder)  {

	return nil
}

func Get (cr *dto.CacheRequest) (r dto.Responder) {
	return nil
}

func Remove (cr *dto.CacheRequest) (r dto.Responder) {
	return nil
}

func Keys () (r dto.Responder) {
	return nil
}

func Capacity() (dto.Responder) {
	cache := *Singltone()
	val := cache.Cap()
	return &dto.CapacityResponse{
		Val: val,
		Err: nil,
	}
}

