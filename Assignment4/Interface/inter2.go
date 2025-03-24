package main

import "fmt"

func process(n int, f func(int) int) int {
    return f(n)
}

func main() {
    result := process(10, func(x int) int {
        return x * x
    })
    fmt.Println("Square:", result)
}
