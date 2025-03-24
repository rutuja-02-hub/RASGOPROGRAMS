package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "Hello, Go!"
    fmt.Println("Contains 'Go':", strings.Contains(str, "Go"))
    fmt.Println("Contains 'Java':", strings.Contains(str, "Java"))
}
