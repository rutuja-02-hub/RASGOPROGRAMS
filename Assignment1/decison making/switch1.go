package main
import "fmt"

func main() {
    var choice int
    var num1, num2 float64
    fmt.Println("Menu:")
    fmt.Println("1. Add")
    fmt.Println("2. Subtract")
    fmt.Println("3. Multiply")
    fmt.Println("4. Divide")
    fmt.Print("Enter your choice: ")
    fmt.Scan(&choice)

    fmt.Print("Enter first number: ")
    fmt.Scan(&num1)
    fmt.Print("Enter second number: ")
    fmt.Scan(&num2)

    switch choice {
    case 1:
        fmt.Println("Sum:", num1 + num2)
    case 2:
        fmt.Println("Difference:", num1 - num2)
    case 3:
        fmt.Println("Product:", num1 * num2)
    case 4:
        if num2 != 0 {
            fmt.Println("Quotient:", num1 / num2)
        } else {
            fmt.Println("Error! Division by zero.")
        }
    default:
        fmt.Println("Invalid choice")
    }
}
