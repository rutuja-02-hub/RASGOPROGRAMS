package main

import "fmt"

func main() {
    var radius float64
    fmt.Print("Enter the radius of the circle: ")
    fmt.Scanln(&radius)
    
    area := 3.14159 * radius * radius
    circumference := 2 * 3.14159 * radius
    
    fmt.Printf("Area of the circle: %.2f\n", area)
    fmt.Printf("Circumference of the circle: %.2f\n", circumference)
}
