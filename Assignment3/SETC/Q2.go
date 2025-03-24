package main

import (
	"fmt"
)

type Employee struct {
	Eno    int
	Ename  string
	Salary float64
}

func main() {
	var n int
	fmt.Print("Enter the number of employees: ")
	fmt.Scan(&n)

	employees := make([]Employee, n)
	var maxSalary float64

	for i := 0; i < n; i++ {
		fmt.Printf("Enter details for employee %d:\n", i+1)
		fmt.Print("Employee Number: ")
		fmt.Scan(&employees[i].Eno)
		fmt.Print("Employee Name: ")
		fmt.Scan(&employees[i].Ename)
		fmt.Print("Salary: ")
		fmt.Scan(&employees[i].Salary)
		if employees[i].Salary > maxSalary {
			maxSalary = employees[i].Salary
		}
	}

	fmt.Println("\nEmployees with Maximum Salary:")
	for _, emp := range employees {
		if emp.Salary == maxSalary {
			fmt.Printf("Eno: %d, Name: %s, Salary: %.2f\n", emp.Eno, emp.Ename, emp.Salary)
		}
	}
}
