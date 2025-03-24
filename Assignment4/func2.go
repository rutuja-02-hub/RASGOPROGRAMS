package main

import "fmt"

func calculate(a int, b int) (int, int) {
    sum := a + b
    product := a * b
    return sum, product
}

func main() {
    sum, product := calculate(5, 10)
    fmt.Println("Sum:", sum)
    fmt.Println("Product:", product)
}
