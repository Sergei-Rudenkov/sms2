package main

import (
	"fmt"
	"sms2/test/client"
	"strconv"
	"sync"
	"time"
	"math/rand"
)

const sms2Output string = "./go_build_main_go"

func main() {
	performanceHttpCacheTest()
}

// run cache first "sms2.exe fixed 3000 60" or "sms2.exe agile"
func performanceHttpCacheTest(){
	var wg1, wg2 sync.WaitGroup
	c := client.HttpClient{Host: "localhost", Port: 8080}

	// start test
	setStart := time.Now()
	for i := 1; i <= 3000; i++ {
		key := strconv.Itoa(i)
		go func() {
			wg1.Add(1)
			defer wg1.Done()
			c.HttpSet("key"+key, "value", 0)
		}()
	}
	wg1.Wait()
	fmt.Printf("Time that took for 3K set operations via http:      %f sec.\n", time.Since(setStart).Seconds())

	getStart := time.Now()
	for i := 1; i <= 3000; i++ {
		go func() {
			wg2.Add(1)
			defer wg2.Done()
			c.HttpGet("key" + strconv.Itoa(rand.Intn(3000)))
		}()
	}
	wg2.Wait()
	fmt.Printf("Time that took for 3K get operations via http:      %f sec.\n", time.Since(getStart).Seconds())
}

//Agile:
//Time that took for 3K set operations via http:      1.733671 sec.
//Time that took for 3K get operations via http:      1.590002 sec.

//Fixed:
//Time that took for 3K set operations via http:      1.430003 sec.
//Time that took for 3K get operations via http:      1.322996 sec.