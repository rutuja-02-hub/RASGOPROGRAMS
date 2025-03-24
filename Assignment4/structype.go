package main

import "fmt"

type Person struct {
    name string
    age  int
}

func (p Person) display() {
    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
}

func main() {
    p1 := Person{name: "Alice", age: 25}
    p1.display()
}
