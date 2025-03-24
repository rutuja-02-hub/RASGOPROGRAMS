package main

import "fmt"

// Declare constants
const Pi float64 = 3.14159
const Greeting = "Hello, Go!"  // The type will be inferred as string

func main() {
    fmt.Println("Value of Pi:", Pi)
    fmt.Println(Greeting)
}
