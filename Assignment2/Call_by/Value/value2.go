package main
import "fmt"

func add(a, b int) int {
    a = a + 1 // Incrementing a local copy of a
    return a + b
}

func main() {
    x, y := 5, 10
    fmt.Println("Before calling add, x =", x, ", y =", y)

    result := add(x, y) // Copies of x and y are passed
    fmt.Println("Result of add:", result)
    fmt.Println("After calling add, x =", x, ", y =", y)
}
