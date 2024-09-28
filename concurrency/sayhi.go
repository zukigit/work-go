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
	string_chanel := make(chan string)
	received := false

	go sayHiAfter(1, chnnel)
	fmt.Println(msg)

	response := <-chnnel
	fmt.Println("response", response)
	response = <-chnnel
	fmt.Println("response", response)

	go func() {
		string_chanel <- "hello mf"
	}()

	for !received {
		select {
		case <-string_chanel:
			fmt.Println("received message")
			received = true
		default:
			fmt.Println("not reeived message")
			time.Sleep(time.Second * 3)
		}
	}
}
