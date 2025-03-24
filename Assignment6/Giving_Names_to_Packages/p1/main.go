package main

import (
	"Assignment6/geometry"
	"fmt"
)

func main() {
	// Create a Rectangle instance
	rect := geometry.Rectangle{Length: 10, Width: 5}

	// Calculate area and perimeter
	area := rect.Area()
	perimeter := rect.Perimeter()

	// Print results
	fmt.Println("Rectangle Area:", area)
	fmt.Println("Rectangle Perimeter:", perimeter)
}
