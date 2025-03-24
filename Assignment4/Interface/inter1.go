package main

import "fmt"

func operate(a, b int, op func(int, int) int) int {
    return op(a, b)
}

func add(x, y int) int {
    return x + y
}

func main() {
    result := operate(5, 3, add)
    fmt.Println("Result:", result)
}
