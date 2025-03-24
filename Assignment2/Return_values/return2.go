package main
import "fmt"

func minMax(a, b int) (int, int) {
    if a < b {
        return a, b
    }
    return b, a
}

func main() {
    x, y := 42, 7
    min, max := minMax(x, y)

    fmt.Println("Minimum:", min)
    fmt.Println("Maximum:", max)
}
