package main

import (
	"fmt"
	"sort"
)

// Define Student struct
type Student struct {
	Name  string
	Marks int
}

// Sorting logic for slice of Students
type ByMarks []Student

func (s ByMarks) Len() int           { return len(s) }
func (s ByMarks) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByMarks) Less(i, j int) bool { return s[i].Marks > s[j].Marks } // Sort in descending order

func main() {
	// Sample student data
	students := []Student{
		{"Alice", 85},
		{"Bob", 92},
		{"Charlie", 78},
		{"David", 90},
	}

	// Sorting students by marks
	sort.Sort(ByMarks(students))

	// Display sorted students
	fmt.Println("Students sorted by marks:")
	for _, student := range students {
		fmt.Println(student.Name, "-", student.Marks)
	}
}
