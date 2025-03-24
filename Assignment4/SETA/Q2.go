package main

import "fmt"

type Numbers struct {
    a, b int
}

func (n Numbers) multiply() int {
    return n.a * n.b
}

func main() {
    num := Numbers{a: 5, b: 10}
    fmt.Println("Multiplication:", num.multiply())
}
