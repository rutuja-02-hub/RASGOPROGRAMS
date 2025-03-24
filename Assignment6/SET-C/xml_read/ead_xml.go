package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Define struct to match XML structure
type Student struct {
	Name  string `xml:"Name"`
	Age   int    `xml:"Age"`
	Marks int    `xml:"Marks"`
}

type Students struct {
	StudentList []Student `xml:"Student"`
}

func main() {
	// Open XML file
	file, err := os.Open("xml_handling/data.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read file content
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal XML into struct
	var students Students
	err = xml.Unmarshal(bytes, &students)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
		return
	}

	// Display structure data
	fmt.Println("Student Data:")
	for _, student := range students.StudentList {
		fmt.Printf("Name: %s, Age: %d, Marks: %d\n", student.Name, student.Age, student.Marks)
	}
}
