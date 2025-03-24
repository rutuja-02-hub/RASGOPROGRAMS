package main

import "fmt"

func main() {
    var arr = [5]int{1, 2, 3, 4, 5}  // Declare and initialize an array
    fmt.Println("Array elements:")
    
    for i := 0; i < len(arr); i++ {  // Iterate through the array and print each element
        fmt.Println(arr[i])
    }
}
