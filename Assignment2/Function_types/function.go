package main
import "fmt"

func operation(a, b int, op func(int, int) int) int {
    return op(a, b)
}

func main() {
    multiply := func(x, y int) int { return x * y }
    result := operation(3, 4, multiply)
    fmt.Println("Multiplication Result:", result)
}
