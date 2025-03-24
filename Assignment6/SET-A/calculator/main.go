package main

import (
	"Assignment6/calculator"
	"fmt"
)

func main() {
	var a, b float64
	var operator string

	// User input
	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	// Perform calculation
	result, err := calculator.Calculate(a, b, operator)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
