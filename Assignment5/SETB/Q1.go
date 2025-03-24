package main

import (
	"fmt"
)

func main() {
	bufferedChan := make(chan int, 5)

	bufferedChan <- 10
	bufferedChan <- 20
	bufferedChan <- 30

	fmt.Printf("Channel capacity: %d\n", cap(bufferedChan))
	fmt.Printf("Channel length before reading: %d\n", len(bufferedChan))

	fmt.Println("Reading values from channel:")
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)

	fmt.Printf("Channel length after reading: %d\n", len(bufferedChan))
}
