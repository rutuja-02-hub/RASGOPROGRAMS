package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice:", slice)

	// Append elements
	slice = append(slice, 6, 7, 8)
	fmt.Println("After Append:", slice)

	// Remove element at index 2
	index := 2
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println("After Removing Element at Index 2:", slice)

	// Copy slice
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	fmt.Println("Copied Slice:", newSlice)

	// Slice Capacity and Length
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))
}