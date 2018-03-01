package storage

import (
	"sms2/storage/dto"
)
var(
	c = Singleton()
)
func Set (cr *dto.CacheRequest) (r dto.Responder)  {
	succ := c.Set(cr.Key, cr.Value)
	return &dto.SetResponse{
		Success: succ,
	}
}

func Get (cr *dto.CacheRequest) (r dto.Responder) {
	value, succ :=c.Get(cr.Key)
	return &dto.GetResponse{
		Value: value, Success:succ,
	}
}

func Remove (cr *dto.CacheRequest) (r dto.Responder) {
	succ := c.Del(cr.Key)
	return &dto.RemoveResponse{
		Success: succ,
	}
}

func Keys () (dto.Responder) {
	return &dto.KeysResponse{
		List: c.Keys(),
	}
}

func Capacity() (dto.Responder) {
	val := c.Cap()
	return &dto.CapacityResponse{
		Val: val,
	}
}

