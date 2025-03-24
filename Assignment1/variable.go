package main

import "fmt"

func main() {
    // Explicit variable declaration
    var age int = 25
    var name string = "Alice"
    var isEmployed bool = true
    
    // Implicit variable declaration
    var height = 5.6  // type inferred as float64
    city := "New York" // type inferred as string

    // Printing the values
    fmt.Println("Age:", age)
    fmt.Println("Name:", name)
    fmt.Println("Employed:", isEmployed)
    fmt.Println("Height:", height)
    fmt.Println("City:", city)
}
