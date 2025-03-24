package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "Hello, Go!"
    fmt.Println("Index of 'Go':", strings.Index(str, "Go"))
    fmt.Println("Index of 'Java':", strings.Index(str, "Java"))
}
