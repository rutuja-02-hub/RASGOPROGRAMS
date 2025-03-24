o
package main

import "fmt"

func process(a, b int, op func(int, int) (int, int)) (int, int) {
    return op(a, b)
}

func sumAndDiff(x, y int) (int, int) {
    return x + y, x - y
}

func main() {
    sum, diff := process(10, 5, sumAndDiff)
    fmt.Println("Sum:", sum, "Difference:", diff)
}
