package main

import (
    "fmt"
    "math"
)

type Shape interface {
    area() float64
    perimeter() float64
}

type Circle struct {
    radius float64
}

type Rectangle struct {
    length, width float64
}

func (c Circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

func (r Rectangle) area() float64 {
    return r.length * r.width
}

func (r Rectangle) perimeter() float64 {
    return 2 * (r.length + r.width)
}

func main() {
    var s Shape

    c := Circle{radius: 5}
    r := Rectangle{length: 10, width: 4}

    s = c
    fmt.Println("Circle Area:", s.area())
    fmt.Println("Circle Perimeter:", s.perimeter())

    s = r
    fmt.Println("Rectangle Area:", s.area())
    fmt.Println("Rectangle Perimeter:", s.perimeter())
}
