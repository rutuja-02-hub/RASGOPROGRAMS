package main

import (
	"fmt"
)

func main() {
	// Creating a buffered channel with a capacity of 2
	ch := make(chan string, 2)

	// Sending values (does not block because the buffer has space)
	ch <- "Message 1"
	ch <- "Message 2"

	// Receiving values
	fmt.Println(<-ch) // Prints "Message 1"
	fmt.Println(<-ch) // Prints "Message 2"
}
