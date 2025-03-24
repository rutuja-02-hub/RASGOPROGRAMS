package main

import "fmt"

// Function to print the Fibonacci series
func printFibonacci(n int) {
    a, b := 0, 1  // Initial two terms of the Fibonacci series
    
    fmt.Print("Fibonacci Series: ")

    // Print the Fibonacci series up to n terms
    for i := 0; i < n; i++ {
        fmt.Print(a, " ")
        // Update a and b for the next terms in the series
        a, b = b, a+b
    }
    fmt.Println() // Move to the next line after printing the series
}

func main() {
    var terms int

    // Input the number of terms for the Fibonacci series
    fmt.Print("Enter the number of terms in the Fibonacci series: ")
    fmt.Scan(&terms)

    // Call the function to print the Fibonacci series
    printFibonacci(terms)
}
