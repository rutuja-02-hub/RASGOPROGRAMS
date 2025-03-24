package main

import "fmt"

// Function to modify the value of variable by passing its pointer
func modifyValue(num *int) {
    *num = 100  // Dereferencing the pointer to modify the original variable
}

func main() {
    a := 50
    fmt.Println("Before:", a)

    modifyValue(&a)  // Passing the address of 'a'
    
    fmt.Println("After:", a)  // Value of 'a' is modified inside the function
}
