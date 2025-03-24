package main

import "fmt"

func main() {
    var complexNum1, complexNum2 complex64
    fmt.Print("Enter the real and imaginary parts of the first complex number (e.g., 2+3i): ")
    fmt.Scanln(&complexNum1)
    fmt.Print("Enter the real and imaginary parts of the second complex number (e.g., 1+4i): ")
    fmt.Scanln(&complexNum2)
    
    sum := complexNum1 + complexNum2
    difference := complexNum1 - complexNum2
    
    fmt.Printf("Sum of the complex numbers: %.2f\n", sum)
    fmt.Printf("Difference of the complex numbers: %.2f\n", difference)
}
