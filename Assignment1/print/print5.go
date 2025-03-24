package main

import "fmt"

func main() {
    var principal, rate, time, interest float64
    fmt.Print("Enter Principal amount: ")
    fmt.Scanln(&principal)
    fmt.Print("Enter Rate of interest: ")
    fmt.Scanln(&rate)
    fmt.Print("Enter Time (in years): ")
    fmt.Scanln(&time)
    
    interest = (principal * rate * time) / 100
    fmt.Printf("The Simple Interest is: %.2f\n", interest)
}
