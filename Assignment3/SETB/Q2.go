package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{10, 5, 8, 3, 2, 7, 4, 1, 9, 6}

	fmt.Println("Original Array:", arr)

	sort.Ints(arr)

	fmt.Println("Sorted Array in Ascending Order:", arr)
}
