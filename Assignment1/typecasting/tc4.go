package main

import (
    "fmt"
    "strconv"
)

func main() {
    var num int
    fmt.Print("Enter an integer: ")
    fmt.Scanln(&num)
    
    // Type casting from int to string using strconv.Itoa
    str := strconv.Itoa(num)
    
    fmt.Printf("The integer %d is converted to string: '%s'\n", num, str)
}
