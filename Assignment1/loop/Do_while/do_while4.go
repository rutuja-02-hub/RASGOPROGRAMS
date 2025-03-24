package main

import "fmt"

func main() {
    var num int
    for {
        fmt.Print("Enter a positive number: ")
        fmt.Scan(&num)
        if num > 0 {
            break
        }
        fmt.Println("Please enter a positive number.")
    }
    fmt.Println("You entered:", num)
}
