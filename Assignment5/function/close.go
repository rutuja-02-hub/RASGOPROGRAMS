package main

import (
	"fmt"
)

func main() {
	// Create a channel
	ch := make(chan int)

	// Sender Goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i // Send values to the channel
		}
		close(ch) // Close the channel after sending values
	}()

	// Receiving values from the channel
	for value := range ch {
		fmt.Println("Received:", value)
	}

	fmt.Println("Channel closed. No more values to receive.")
}
