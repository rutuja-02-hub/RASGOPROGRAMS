package main
import "fmt"

func divide(a, b int) (int, int) {
    return a / b, a % b
}

func main() {
    quotient, remainder := divide(10, 3)
    fmt.Println("Quotient:", quotient, "Remainder:", remainder)
}
