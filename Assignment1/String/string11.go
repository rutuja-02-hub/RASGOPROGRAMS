package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "Hello, Go!"
    fmt.Println("Starts with 'Hello':", strings.HasPrefix(str, "Hello"))
    fmt.Println("Starts with 'Go':", strings.HasPrefix(str, "Go"))
}
