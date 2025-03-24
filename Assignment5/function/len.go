package main

import (
	"fmt"
)

func main() {
	// Create a buffered channel with a capacity of 3
	ch := make(chan int, 3)

	// Sending values into the channel
	ch <- 10
	ch <- 20

	// Checking length and capacity
	fmt.Println("Current length of channel:", len(ch)) // Output: 2
	fmt.Println("Capacity of channel:", cap(ch))       // Output: 3

	// Receiving one value
	<-ch

	fmt.Println("Length after receiving one value:", len(ch)) // Output: 1
}
