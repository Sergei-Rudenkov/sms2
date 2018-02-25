package main

import (
	"fmt"
	"sms2/storage"
	"sms2/server"
)


func main ()  {
	fmt.Println("Hello Go!")
	s := storage.New(100, storage.WithTTL(10 * 10000))
	s.Set("first_value", "Storage is working")
	test, _ := s.Get("first_value")
	fmt.Println(test)
	server.Start()
}
