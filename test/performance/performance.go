package performance

import (
	"fmt"
	//log "github.com/inconshreveable/log15"
	"sms2/test/client"
	"sync"
	"strconv"
)

const sms2Output string = "./go_build_main_go"

func main() {
	httpFixedCacheTest()
	httpAgileCacheTest()
}

func httpFixedCacheTest(){
	// run cache first "sms2.exe fixed 1000 60"
	var wg sync.WaitGroup

	// start test
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		go func(){
			defer wg.Done()
			key := strconv.Itoa(i)
			client.HttpFixedSet(key, "value")
		}()
	}
	wg.Wait()
}

func httpAgileCacheTest(){
	fmt.Println("httpAgileCacheTest")
}

