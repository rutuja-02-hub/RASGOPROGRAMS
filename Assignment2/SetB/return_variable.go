package main
import "fmt"

func calculate(a, b int) (sum int, product int) {
    sum = a + b
    product = a * b
    return
}

func main() {
    num1, num2 := 5, 10
    sum, product := calculate(num1, num2)
    fmt.Println("Sum:", sum)
    fmt.Println("Product:", product)
}
