package main

import "fmt"

func main() {
    var name string
    var age int
    
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name) // Taking string input from user
    
    fmt.Print("Enter your age: ")
    fmt.Scanln(&age) // Taking integer input from user
    
    fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}
