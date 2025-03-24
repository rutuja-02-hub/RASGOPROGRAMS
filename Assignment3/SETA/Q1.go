package main

import (
	"fmt"
)

func findMinMax(arr []int) (int, int) {
	if len(arr) == 0 {
		panic("Array is empty")
	}

	smallest, largest := arr[0], arr[0]

	for _, num := range arr {
		if num < smallest {
			smallest = num
		}
		if num > largest {
			largest = num
		}
	}

	return smallest, largest
}

func main() {
	arr := []int{10, 2, 8, 3, 6, 1, 9, 4}
	smallest, largest := findMinMax(arr)
	fmt.Printf("Smallest number: %d\n", smallest)
	fmt.Printf("Largest number: %d\n", largest)
}
