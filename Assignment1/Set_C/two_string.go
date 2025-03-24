package main

import "fmt"

func main() {
    // Using new() to create a pointer to an integer
    ptr := new(int)  // Allocates memory for an int and returns a pointer to it

    // Print the value and address of the variable
    fmt.Println("Value of ptr (initially):", *ptr)  // Dereferencing to get the value (default is 0)
    fmt.Println("Address of ptr:", ptr)

    // Assign a value to the memory location
    *ptr = 42

    // Print the updated value
    fmt.Println("Updated value of ptr:", *ptr)  // Dereferencing again to get the updated value
}
