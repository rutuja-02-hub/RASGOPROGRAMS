package main

import "fmt"

type Author struct {
    name  string
    books int
}

func (a Author) display() {
    fmt.Println("Author Name:", a.name)
    fmt.Println("Number of Books:", a.books)
}

func main() {
    author1 := Author{name: "J.K. Rowling", books: 7}
    author1.display()
}
