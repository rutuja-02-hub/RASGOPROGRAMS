package main

import "fmt"

func applyOperation(a, b int, op func(int, int) int) int {
    return op(a, b)
}

func main() {
    fmt.Println("Addition:", applyOperation(8, 3, func(x, y int) int { return x + y }))
    fmt.Println("Subtraction:", applyOperation(8, 3, func(x, y int) int { return x - y }))
}
