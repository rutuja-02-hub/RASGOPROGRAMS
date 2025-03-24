package main

import "fmt"

type Person struct {
    name string
    age  int
}

// Value Receiver (Does not modify original)
func (p Person) show() {
    fmt.Println("Name:", p.name, "Age:", p.age)
}

// Pointer Receiver (Modifies original)
func (p *Person) updateAge(newAge int) {
    p.age = newAge
}

func main() {
    p1 := Person{name: "Alice", age: 25}

    p1.show()        // No modification
    p1.updateAge(30) // Modifies the original struct
    p1.show()        // Updated value reflected
}
