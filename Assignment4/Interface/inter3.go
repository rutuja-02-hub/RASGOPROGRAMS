package main

import "fmt"

type Operation func(int, int) int

func calculate(a, b int, op Operation) int {
    return op(a, b)
}

func multiply(x, y int) int {
    return x * y
}

func main() {
    result := calculate(4, 5, multiply)
    fmt.Println("Multiplication:", result)
}
