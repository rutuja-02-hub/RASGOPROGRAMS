package main

import (
	"Assignment6/rectangle"
	"fmt"
)

func main() {
	var length, width float64

	// User input
	fmt.Print("Enter length of the rectangle: ")
	fmt.Scanln(&length)
	fmt.Print("Enter width of the rectangle: ")
	fmt.Scanln(&width)

	// Calculate and display area
	area := rectangle.CalculateArea(length, width)
	fmt.Println("Area of the rectangle:", area)
}
