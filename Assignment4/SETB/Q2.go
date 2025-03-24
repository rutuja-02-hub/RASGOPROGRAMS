package main

import "fmt"

func checkType(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Type: int, Value:", v)
    case float64:
        fmt.Println("Type: float64, Value:", v)
    case string:
        fmt.Println("Type: string, Value:", v)
    case bool:
        fmt.Println("Type: bool, Value:", v)
    default:
        fmt.Println("Unknown type")
    }
}

func main() {
    checkType(10)
    checkType(3.14)
    checkType("Hello")
    checkType(true)
}
