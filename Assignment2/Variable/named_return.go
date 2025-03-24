package main
import "fmt"

func calculate(a, b int) (sum int, product int) {
    
    sum = a + b
    product = a * b
    return 
}

func main() {
    x, y := 5, 10
    sum, product := calculate(x, y)

    fmt.Println("Sum:", sum)
    fmt.Println("Product:", product)
}

