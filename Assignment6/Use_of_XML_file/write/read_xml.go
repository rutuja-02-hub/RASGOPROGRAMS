package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Define struct for student data
type Student struct {
	Name  string `xml:"Name"`
	Age   int    `xml:"Age"`
	Marks int    `xml:"Marks"`
}

type Students struct {
	StudentList []Student `xml:"Student"`
}

func main() {
	// Open the XML file
	file, err := os.Open("xml_example/students.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file content
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal XML into struct
	var students Students
	err = xml.Unmarshal(bytes, &students)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return
	}

	// Print student data
	fmt.Println("Student Data:")
	for _, student := range students.StudentList {
		fmt.Printf("Name: %s, Age: %d, Marks: %d\n", student.Name, student.Age, student.Marks)
	}
}
