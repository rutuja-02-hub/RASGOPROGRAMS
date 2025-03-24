package main

import (
	"fmt"
)

func main() {
	// Creating an unbuffered channel
	ch := make(chan string)

	// Goroutine to send data
	go func() {
		ch <- "Hello from Goroutine!" // This will block until received
	}()

	// Receiving data
	msg := <-ch
	fmt.Println("Received:", msg)
}
