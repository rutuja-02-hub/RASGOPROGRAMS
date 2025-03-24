package main

import "fmt"

func main() {
    arr := [6]int{34, 56, 12, 89, 23, 67}
    largest := arr[0]

    for i := 1; i < len(arr); i++ {
        if arr[i] > largest {
            largest = arr[i]  // Update the largest if a larger element is found
        }
    }

    fmt.Println("Largest element in the array:", largest)
}
