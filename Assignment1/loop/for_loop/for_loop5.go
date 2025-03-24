package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    
    factorial := 1
    for i := 1; i <= num; i++ {
        factorial *= i
    }
    fmt.Println("Factorial of", num, "is", factorial)
}
