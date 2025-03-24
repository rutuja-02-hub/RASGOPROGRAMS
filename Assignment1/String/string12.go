package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "Hello, Go!"
    fmt.Println("Ends with 'Go!':", strings.HasSuffix(str, "Go!"))
    fmt.Println("Ends with 'Hello':", strings.HasSuffix(str, "Hello"))
}
