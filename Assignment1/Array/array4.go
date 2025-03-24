package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}

    fmt.Println("Original array:", arr)

    // Reversing the array
    for i := 0; i < len(arr)/2; i++ {
        arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]  // Swap elements
    }

    fmt.Println("Reversed array:", arr)
}
