package main

import (
	"fmt"
)

func main() {
	// Create a buffered channel with a capacity of 3
	ch := make(chan int, 3)

	fmt.Println("Capacity of channel:", cap(ch)) // Output: 3

	// Sending values
	ch <- 10
	ch <- 20

	// Checking length (current items) and capacity
	fmt.Println("Current length of channel:", len(ch)) // Output: 2
	fmt.Println("Capacity of channel:", cap(ch))       // Output: 3
}
