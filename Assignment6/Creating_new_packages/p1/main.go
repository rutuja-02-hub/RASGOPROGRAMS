package main

import (
	"Assignment6/mathutils"
	"fmt"
)

func main() {
	a, b := 10, 5

	// Using mathutils package
	sum := mathutils.Add(a, b)
	product := mathutils.Multiply(a, b)

	// Print results
	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)
}
