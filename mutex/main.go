package main

import (
	"fmt"
	"sync"
	"time"
)

var v = make(map[string]int)
var mu sync.Mutex

func adddOne() {
	mu.Lock()
	v["test"]++
	mu.Unlock()
}

func main() {
	fmt.Println("It's mutex")

	for i := 0; i < 100; i++ {
		go adddOne()
	}

	time.Sleep(time.Second)
	fmt.Println("v['test']", v["test"])

	webCrawler()
}
