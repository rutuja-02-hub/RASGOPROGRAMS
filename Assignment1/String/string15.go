package main

import "fmt"

func main() {
    str := "Hello, Go!"
    runes := []rune(str)

    fmt.Println("Rune slice:", runes)
    fmt.Println("First character:", string(runes[0]))
    fmt.Println("Last character:", string(runes[len(runes)-1]))
}
