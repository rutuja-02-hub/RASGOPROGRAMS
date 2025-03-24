package main

import (
	"fmt"
)

func evenNumbers(evenChan chan int) {
	for num := range evenChan {
		fmt.Printf("Even: %d\n", num)
	}
}

func oddNumbers(oddChan chan int) {
	for num := range oddChan {
		fmt.Printf("Odd: %d\n", num)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenChan := make(chan int)
	oddChan := make(chan int)

	go evenNumbers(evenChan)
	go oddNumbers(oddChan)

	for _, num := range numbers {
		if num%2 == 0 {
			evenChan <- num
		} else {
			oddChan <- num
		}
	}

	close(evenChan)
	close(oddChan)
}
