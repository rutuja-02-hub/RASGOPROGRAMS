package main

import "fmt"

type Person struct {
    name string
    age  int
}

// Method with value receiver
func (p Person) show() {
    fmt.Println("Name:", p.name, "Age:", p.age)
}

// Method with pointer receiver
func (p *Person) updateAge(newAge int) {
    p.age = newAge
}

func main() {
    p1 := Person{name: "Alice", age: 25}

    p1.show()       // Value receiver (does not modify original)
    p1.updateAge(30) // Pointer receiver (modifies original)
    p1.show()       // Updated age reflected
}
