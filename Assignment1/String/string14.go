package main

import "fmt"

func main() {
    name := "Go"
    version := 1.20
    formattedString := fmt.Sprintf("Programming Language: %s, Version: %.2f", name, version)
    fmt.Println(formattedString)
}
