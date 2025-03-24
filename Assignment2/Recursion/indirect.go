package main
import "fmt"

func isEven(n int) bool {
    if n == 0 {
        return true 
    }
    return isOdd(n - 1) 
}

func isOdd(n int) bool {
    if n == 0 {
        return false 
    }
    return isEven(n - 1) 
}

func main() {
    num := 7
    if isEven(num) {
        fmt.Println(num, "is Even")
    } else {
        fmt.Println(num, "is Odd")
    }
}
