package main

import (
	"fmt"
)

func main() {
	// Create an unbuffered channel of type int
	ch := make(chan int)

	// Goroutine to send data
	go func() {
		fmt.Println("Sending data...")
		ch <- 10 // Send value 10 into the channel
	}()

	// Receiving data
	value := <-ch
	fmt.Println("Received:", value)
}
