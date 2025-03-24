package main

import "fmt"

// Define a struct for Student
type Student struct {
    Name     string
    RollNo   int
    Division string
    College  string
}

func main() {
    // Create an instance of the Student struct and assign values
    student := Student{
        Name:     "John Doe",
        RollNo:   12345,
        Division: "A",
        College:  "XYZ College",
    }

    // Print student details
    fmt.Println("Student Name:", student.Name)
    fmt.Println("Roll Number:", student.RollNo)
    fmt.Println("Division:", student.Division)
    fmt.Println("College Name:", student.College)
}
