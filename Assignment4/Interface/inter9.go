package main

import "fmt"

func multiplier(factor int) func(int) int {
    return func(n int) int {
        return n * factor
    }
}

func main() {
    double := multiplier(2)
    triple := multiplier(3)

    fmt.Println("Double of 4:", double(4))
    fmt.Println("Triple of 4:", triple(4))
}
