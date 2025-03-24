package main

import "fmt"

type Author struct {
    name  string
    books int
}

func (a Author) show() {
    fmt.Println("Author Name:", a.name)
    fmt.Println("Number of Books:", a.books)
}

func main() {
    author1 := Author{name: "George Orwell", books: 5}
    author1.show()
}
