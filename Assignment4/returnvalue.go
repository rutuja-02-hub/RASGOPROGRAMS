package main

import "fmt"

func calculate(a int, b int) (sum int, product int) {
    sum = a + b
    product = a * b
    return
}

func main() {
    s, p := calculate(4, 7)
    fmt.Println("Sum:", s)
    fmt.Println("Product:", p)
}
