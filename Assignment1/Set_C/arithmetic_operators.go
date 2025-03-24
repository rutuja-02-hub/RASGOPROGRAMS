package main

import "fmt"

func main() {
    var num1, num2 float64
    var choice int

    // Accept two numbers from the user
    fmt.Print("Enter first number: ")
    fmt.Scan(&num1)
    fmt.Print("Enter second number: ")
    fmt.Scan(&num2)

    // Display arithmetic options to the user
    fmt.Println("\nChoose an operation:")
    fmt.Println("1. Addition")
    fmt.Println("2. Subtraction")
    fmt.Println("3. Multiplication")
    fmt.Println("4. Division")

    // Accept the user's choice
    fmt.Print("\nEnter your choice (1/2/3/4): ")
    fmt.Scan(&choice)

    // Perform the operation based on user's choice
    switch choice {
    case 1:
        result := num1 + num2
        fmt.Printf("\n%.2f + %.2f = %.2f\n", num1, num2, result)
    case 2:
        result := num1 - num2
        fmt.Printf("\n%.2f - %.2f = %.2f\n", num1, num2, result)
    case 3:
        result := num1 * num2
        fmt.Printf("\n%.2f * %.2f = %.2f\n", num1, num2, result)
    case 4:
        if num2 != 0 {
            result := num1 / num2
            fmt.Printf("\n%.2f / %.2f = %.2f\n", num1, num2, result)
        } else {
            fmt.Println("\nError: Division by zero is not allowed.")
        }
    default:
        fmt.Println("\nInvalid choice! Please enter a number between 1 and 4.")
    }
}
