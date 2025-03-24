package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "Hello, Go!"
    newStr := strings.Replace(str, "Go", "Golang", 1)
    fmt.Println("Replaced string:", newStr)
}
