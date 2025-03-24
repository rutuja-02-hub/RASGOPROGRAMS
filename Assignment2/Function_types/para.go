package main
import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    sum := add(5, 7)
    fmt.Println("Sum:", sum)
}
