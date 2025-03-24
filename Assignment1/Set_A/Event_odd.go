package main

import "fmt"

func main() {
    // Declare a variable to hold the number
    var num int
    
    // Ask the user for input
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)  // Take input from the user

    // Check if the number is even or odd
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }
}
