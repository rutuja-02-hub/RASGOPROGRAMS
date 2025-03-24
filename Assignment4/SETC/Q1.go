package main

import "fmt"

type Info interface{}

func displayValue(i Info) {
    value, ok := i.(string)
    if ok {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Type assertion failed")
    }
}

func main() {
    var data Info = "Hello, Go!"
    displayValue(data)
}
