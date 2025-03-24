package main
import "fmt"

func swap(a, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

func main() {
    x, y := 5, 10
    fmt.Println("Before swap: x =", x, ", y =", y)

    swap(&x, &y) 
    fmt.Println("After swap: x =", x, ", y =", y)
}
