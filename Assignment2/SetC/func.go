package main
import "fmt"

func calculate(a, b int) (int, int, int) {
    sum := a + b
    diff := a - b
    product := a * b
    return sum, diff, product
}

func main() {
    num1, num2 := 10, 5

    sum, diff, product := calculate(num1, num2)

    fmt.Println("Sum:", sum)
    fmt.Println("Difference:", diff)
    fmt.Println("Product:", product)
}
