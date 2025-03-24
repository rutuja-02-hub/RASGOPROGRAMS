package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    sum := add(5, 10)
    fmt.Println("Sum:", sum)

    greet("Alice")
}
