package main
import "fmt"

func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

func main() {
    x, y := 10, 3
    quotient, remainder := divide(x, y)

    fmt.Println("Quotient:", quotient)
    fmt.Println("Remainder:", remainder)
}
