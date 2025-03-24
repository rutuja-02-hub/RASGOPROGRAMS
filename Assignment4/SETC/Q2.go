package main

import (
    "fmt"
    "sort"
)

type Student struct {
    rollNo     int
    name       string
    percentage float64
}

type StudentList struct {
    students []Student
}

func (s *StudentList) addStudent(rollNo int, name string, percentage float64) {
    s.students = append(s.students, Student{rollNo, name, percentage})
}

func (s *StudentList) displayDescending() {
    sort.Slice(s.students, func(i, j int) bool {
        return s.students[i].percentage > s.students[j].percentage
    })

    fmt.Println("Student Information in Descending Order of Percentage:")
    for _, student := range s.students {
        fmt.Printf("Roll No: %d, Name: %s, Percentage: %.2f%%\n", student.rollNo, student.name, student.percentage)
    }
}

func main() {
    var s StudentList

    s.addStudent(101, "Alice", 85.5)
    s.addStudent(102, "Bob", 92.3)
    s.addStudent(103, "Charlie", 78.9)
    s.addStudent(104, "David", 88.1)

    s.displayDescending()
}
