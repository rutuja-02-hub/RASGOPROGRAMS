package main

import "fmt"

func main() {
    var a int = 58
    var ptr *int = &a  // Pointer to the address of variable 'a'

    fmt.Println("Value of a:", a)
    fmt.Println("Address of a:", &a)
    fmt.Println("Value stored at ptr:", *ptr) // Dereferencing the pointer
    fmt.Println("Address stored in ptr:", ptr)
}
