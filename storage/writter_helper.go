package storage

import (
	"sms2/storage/dto"
	"fmt"
)
var(
	c = Singleton()
)
func Set (cr *dto.CacheRequest) (r dto.Responder)  {
	updated := c.Set(cr.Key, cr.Value)
	//fmt.Println(c.Keys())
	return &dto.SetResponse{
		Evicted: updated,
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

