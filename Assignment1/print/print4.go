package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scanln(&num)
    
    if num%2 == 0 {
        fmt.Println("The number is Even.")
    } else {
        fmt.Println("The number is Odd.")
    }
}
