package main

import "fmt"

func main() {
    var num1 float64
    fmt.Print("Enter a float number: ")
    fmt.Scanln(&num1)
    
    // Type casting from float64 to int (truncates decimal part)
    var num2 int = int(num1)
    
    fmt.Printf("The float %.2f is converted to integer: %d\n", num1, num2)
}
