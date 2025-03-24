package main

import "fmt"

func printSlice(slice []int) {
	fmt.Println("Slice inside function:", slice)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	printSlice(numbers)
}
