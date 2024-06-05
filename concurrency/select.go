package main

import (
	"fmt"
	"time"
)

func sleep_t(signal_chan chan int, sec time.Duration) {
	signal := 0
	for {
		select {
		case signal = <-signal_chan:
			if signal == 69 {
				return
			}
		default:
			fmt.Println("I'm sleeping")
			time.Sleep(sec * time.Second)
		}
	}
}

func selectGo() {
	signal_chan := make(chan int)

	go sleep_t(signal_chan, 10)

	time.Sleep(30 * time.Second)
	signal_chan <- 69
}
