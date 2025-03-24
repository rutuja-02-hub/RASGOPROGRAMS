package main

import "fmt"

func main() {
    var num1 int
    fmt.Print("Enter an integer: ")
    fmt.Scanln(&num1)
    
    // Type casting from int to float64
    var num2 float64 = float64(num1)
    
    fmt.Printf("The integer %d is converted to float: %.2f\n", num1, num2)
}
