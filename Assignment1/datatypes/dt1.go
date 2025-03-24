package main

import "fmt"

func main() {
    var num1, num2 int
    fmt.Print("Enter the first integer: ")
    fmt.Scanln(&num1)
    fmt.Print("Enter the second integer: ")
    fmt.Scanln(&num2)
    
    sum := num1 + num2
    difference := num1 - num2
    product := num1 * num2
    quotient := float64(num1) / float64(num2)

    fmt.Println("Sum:", sum)
    fmt.Println("Difference:", difference)
    fmt.Println("Product:", product)
    fmt.Printf("Quotient: %.2f\n", quotient)
}
