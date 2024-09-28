package main

import (
	"fmt"
	"time"
)

func sayHiAfter(sec time.Duration, channel chan int) {
	time.Sleep(sec * time.Microsecond)
	fmt.Println("HI! after", sec)

	channel <- 100
	channel <- 69
	fmt.Println("send 69")
	close(channel)
}

func sayHi(msg string) {
	chnnel := make(chan int)
	go sayHiAfter(1, chnnel)
	fmt.Println(msg)

	response := <-chnnel
	fmt.Println("response", response)
	// response = <-chnnel
	// fmt.Println("response", response)
}
