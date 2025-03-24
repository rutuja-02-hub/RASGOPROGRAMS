package main

import "fmt"

func main() {
    var a int = 25
    var ptr1 *int = &a       // Pointer to a
    var ptr2 **int = &ptr1   // Pointer to pointer

    fmt.Println("Value of a:", a)
    fmt.Println("Address of a:", &a)
    fmt.Println("Value stored at ptr1:", *ptr1)  // Dereferencing ptr1
    fmt.Println("Address stored in ptr1:", ptr1)
    fmt.Println("Value stored at ptr2:", **ptr2)  // Dereferencing ptr2 twice
    fmt.Println("Address stored in ptr2:", ptr2)
}
