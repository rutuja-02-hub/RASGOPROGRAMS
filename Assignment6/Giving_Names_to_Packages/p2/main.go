package main

import (
	"Assignment6/operations"
	"fmt"
)

func main() {
	a, b := 12, 8

	// Using operations package
	sum := operations.Add(a, b)
	difference := operations.Subtract(a, b)
	product := operations.Multiply(a, b)

	// Print results
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
}
