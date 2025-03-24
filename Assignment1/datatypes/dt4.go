package main

import "fmt"

func main() {
    var isRainy bool
    fmt.Print("Is it raining today? (true/false): ")
    fmt.Scanln(&isRainy)
    
    if isRainy {
        fmt.Println("Don't forget to take an umbrella!")
    } else {
        fmt.Println("It's a nice day!")
    }
}
