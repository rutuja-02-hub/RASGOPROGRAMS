package main

import "fmt"

func main() {
    arr := [5]int{10, 20, 30, 40, 50}
    sum := 0

    for i := 0; i < len(arr); i++ {
        sum += arr[i]  // Add each element to the sum
    }

    fmt.Println("Sum of array elements:", sum)
}
