package main

import (
	"fmt"
	"sms2/storage"
	"sms2/server"
)


func main ()  {
	fmt.Println("Hello Go!")
	server.Start()
}
