package main

import "fmt"

// Function to modify the first element of an array using a pointer
func modifyArray(arr *[5]int) {
    arr[0] = 10  // Modify the first element of the array
}

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Before modification:", arr)

    modifyArray(&arr)  // Passing the address of the array

    fmt.Println("After modification:", arr)  // The first element of the array is modified
}
