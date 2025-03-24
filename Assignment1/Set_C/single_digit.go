package main

import "fmt"

func main() {
    var num int

    // Accept the number from the user
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    // Check if the number is a single digit (between -9 and 9)
    if num >= -9 && num <= 9 {
        fmt.Println("The number is a single digit.")
    } else {
        fmt.Println("The number is not a single digit.")
    }
}
