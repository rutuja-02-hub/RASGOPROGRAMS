package main

import (
    "fmt"
    "strconv"
)

func main() {
    var num float64
    fmt.Print("Enter a float number: ")
    fmt.Scanln(&num)
    
    // Type casting from float to string using strconv.FormatFloat
    str := strconv.FormatFloat(num, 'f', 2, 64)
    
    fmt.Printf("The float %.2f is converted to string: '%s'\n", num, str)
}
