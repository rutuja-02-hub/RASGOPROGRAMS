package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "Hello, Go! How are you?"
    substrings := strings.Split(str, " ")
    fmt.Println("Splitted substrings:", substrings)
}
