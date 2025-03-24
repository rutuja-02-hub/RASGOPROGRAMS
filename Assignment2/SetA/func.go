package main
import "fmt"

func addNumbers(a int, b int) int {
    return a + b
}

func main() {
    var num1, num2 int

    fmt.Print("Enter the first number: ")
    fmt.Scan(&num1)
    fmt.Print("Enter the second number: ")
    fmt.Scan(&num2)

    sum := addNumbers(num1, num2)
    fmt.Println("The sum of", num1, "and", num2, "is:", sum)
}
