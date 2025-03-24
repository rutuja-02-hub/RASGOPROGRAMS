package main

import "fmt"

func main() {
    arr := [6]int{10, 20, 30, 40, 50, 60}
    target := 40
    found := false

    for i := 0; i < len(arr); i++ {
        if arr[i] == target {
            found = true
            break
        }
    }

    if found {
        fmt.Printf("The number %d is present in the array.\n", target)
    } else {
        fmt.Printf("The number %d is not present in the array.\n", target)
    }
}
