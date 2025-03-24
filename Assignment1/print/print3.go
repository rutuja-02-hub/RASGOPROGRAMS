package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scanln(&num)
    
    fmt.Println("Multiplication table for:", num)
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d x %d = %d\n", num, i, num*i)
    }
}
