package main

import "fmt"

// Define a struct
type Person struct {
    Name string
    Age  int
}

// Function to modify struct values using pointers
func modifyPerson(p *Person) {
    p.Name = "Alice"  // Modifying struct fields through pointer
    p.Age = 30
}

func main() {
    person := Person{"Bob", 25}
    fmt.Println("Before modification:", person)

    modifyPerson(&person)  // Passing the address of 'person' to modify it

    fmt.Println("After modification:", person)  // Struct values are modified
}
