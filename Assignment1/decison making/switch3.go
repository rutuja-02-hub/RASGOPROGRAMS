package main
import "fmt"

func main() {
    var choice int
    fmt.Println("Menu:")
    fmt.Println("1. Area of Circle")
    fmt.Println("2. Area of Rectangle")
    fmt.Println("3. Area of Triangle")
    fmt.Print("Enter your choice: ")
    fmt.Scan(&choice)

    switch choice {
    case 1:
        var radius float64
        fmt.Print("Enter the radius of the circle: ")
        fmt.Scan(&radius)
        area := 3.14 * radius * radius
        fmt.Printf("Area of Circle: %.2f\n", area)
    case 2:
        var length, width float64
        fmt.Print("Enter length and width of the rectangle: ")
        fmt.Scan(&length, &width)
        area := length * width
        fmt.Printf("Area of Rectangle: %.2f\n", area)
    case 3:
        var base, height float64
        fmt.Print("Enter base and height of the triangle: ")
        fmt.Scan(&base, &height)
        area := 0.5 * base * height
        fmt.Printf("Area of Triangle: %.2f\n", area)
    default:
        fmt.Println("Invalid choice")
    }
}

