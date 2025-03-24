package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "Go Go Golang Go"
    fmt.Println("Count of 'Go':", strings.Count(str, "Go"))
}
