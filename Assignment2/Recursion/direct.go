package main
import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1 
    }
    return n * factorial(n-1) 
}

func main() {
    num := 5
    result := factorial(num)

    fmt.Println("Factorial of", num, "is", result)
}
