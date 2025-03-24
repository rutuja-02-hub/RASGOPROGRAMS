package main

import (
	"fmt"
)

func main() {
	// Create a new channel of type int
	ch := make(chan int)

	// Launch a goroutine that sends a value into the channel
	go func() {
		ch <- 42 // Send value into the channel
	}()

	// Receive the value from the channel
	value := <-ch
	fmt.Println("Received:", value)
}
