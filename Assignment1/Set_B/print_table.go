package main

import "fmt"

func main() {
    var num int

    // Take input for the number
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    // Print the multiplication table
    fmt.Println("Multiplication table of", num)
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d x %d = %d\n", num, i, num*i)
    }
}
