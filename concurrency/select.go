package main

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomInt(min, max int) int {
	// Use NewSource to create a local random source
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	// Generate a random number between min (inclusive) and max (exclusive)
	return generator.Intn(max-min) + min
}

func GenerateRandomString(length int) string {
	// Use NewSource to create a local random source
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	b := make([]byte, length)
	for i := range b {
		// Generate a random index within the charset string
		index := generator.Intn(len(charset))
		// Assign the character at the random index to the slice
		b[i] = charset[index]
	}
	return string(b)
}

func generateRandom(string_chan chan string, int_chan chan int) {
	randomString := GenerateRandomString(10)
	randomInt := GenerateRandomInt(0, 100)

	for {
		select {
		case string_chan <- randomString:
			randomString = GenerateRandomString(10)
		case int_chan <- randomInt:
			randomInt = GenerateRandomInt(0, 100)
		}
	}
}

func selectGo() {
	string_chan := make(chan string)
	int_chan := make(chan int)

	go func() {
		fmt.Println(<-string_chan)
		fmt.Println(<-int_chan)
	}()
	generateRandom(string_chan, int_chan)
}
