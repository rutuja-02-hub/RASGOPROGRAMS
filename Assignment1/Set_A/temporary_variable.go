package main

import "fmt"

func main() {
    var a, b int

    // Take input for both numbers
    fmt.Print("Enter first number: ")
    fmt.Scan(&a)
    fmt.Print("Enter second number: ")
    fmt.Scan(&b)

    // Print original values
    fmt.Println("Before swapping: a =", a, "b =", b)

    // Swap the numbers without using a temporary variable
    a = a + b
    b = a - b
    a = a - b

    // Print swapped values
    fmt.Println("After swapping: a =", a, "b =", b)
}
