package main

import "fmt"

type Student struct {
    name  string
    grade int
}

func (s *Student) show() {
    fmt.Println("Student Name:", s.name)
    fmt.Println("Grade:", s.grade)
}

func main() {
    student1 := Student{name: "John", grade: 90}
    student1.show()
}
