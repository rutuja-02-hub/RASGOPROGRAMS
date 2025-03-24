package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter an integer: ")
    fmt.Scanln(&num) // Taking integer input from user
    
    fmt.Println("You entered:", num)
}
