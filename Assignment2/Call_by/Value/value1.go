package main
import "fmt"

func modifyValue(num int) {
    num = 100
    fmt.Println("Inside modifyValue, num =", num)
}

func main() {
    x := 42
    fmt.Println("Before calling modifyValue, x =", x)

    modifyValue(x) // A copy of x is passed
    fmt.Println("After calling modifyValue, x =", x)
}
