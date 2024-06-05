package main

import (
	"context"
	"fmt"
	"time"
)

func sleep_t(ctx context.Context, sec time.Duration) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("I'm sleeping")
			time.Sleep(sec * time.Second)
		}
	}
}

func selectGo() {
	ctx, cancel := context.WithCancel(context.Background())

	go sleep_t(ctx, 10)

	time.Sleep(5 * time.Second)
	cancel()
}
