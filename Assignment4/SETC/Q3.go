package main

import "fmt"

type Shape interface {
    area() float64
}

type Perimeter interface {
    perimeter() float64
}

type Geometry interface {
    Shape
    Perimeter
}

type Rectangle struct {
    length, width float64
}

func (r Rectangle) area() float64 {
    return r.length * r.width
}

func (r Rectangle) perimeter() float64 {
    return 2 * (r.length + r.width)
}

func main() {
    var g Geometry = Rectangle{length: 10, width: 5}

    fmt.Println("Area:", g.area())
    fmt.Println("Perimeter:", g.perimeter())
}
