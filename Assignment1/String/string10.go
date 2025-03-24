package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "   Hello, Go!   "
    trimmedStr := strings.TrimSpace(str)
    fmt.Println("Trimmed string:", trimmedStr)
}
