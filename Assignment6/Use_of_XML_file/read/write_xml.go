package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Define struct for student data
type Student struct {
	XMLName xml.Name `xml:"Student"`
	Name    string   `xml:"Name"`
	Age     int      `xml:"Age"`
	Marks   int      `xml:"Marks"`
}

type Students struct {
	XMLName    xml.Name  `xml:"Students"`
	StudentList []Student `xml:"Student"`
}

func main() {
	// Sample student data
	students := Students{
		StudentList: []Student{
			{Name: "Alice", Age: 21, Marks: 85},
			{Name: "Bob", Age: 22, Marks: 90},
		},
	}

	// Create XML file
	file, err := os.Create("xml_example/students.xml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Encode struct to XML and write to file
	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	if err := encoder.Encode(students); err != nil {
		fmt.Println("Error encoding XML:", err)
		return
	}

	fmt.Println("Data successfully written to students.xml")
}
