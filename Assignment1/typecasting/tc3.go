package main

import (
    "fmt"
    "strconv"
)

func main() {
    var str string
    fmt.Print("Enter a number as a string: ")
    fmt.Scanln(&str)
    
    // Type casting string to int using strconv.Atoi
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("The string '%s' is converted to integer: %d\n", str, num)
    }
}
