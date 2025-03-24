package main

import "fmt"

type Person struct {
        FirstName string
        LastName  string
        Age       int
        Address   Address
}

type Address struct {
        Street  string
        City    string
        ZipCode string
}

type Product struct {
        ID    int     `json:"id"`
        Name  string  `json:"name"`
        Price float64 `json:"price"`
}

func main() {
        person1 := Person{
                FirstName: "John",
                LastName:  "Doe",
                Age:       30,
                Address: Address{
                        Street:  "123 Main St",
                        City:    "Anytown",
                        ZipCode: "12345",
                },
        }

        fmt.Println("Person 1:")
        fmt.Println("First Name:", person1.FirstName)
        fmt.Println("Last Name:", person1.LastName)
        fmt.Println("Age:", person1.Age)
        fmt.Println("Street:", person1.Address.Street)
        fmt.Println("City:", person1.Address.City)
        fmt.Println("Zip Code:", person1.Address.ZipCode)

        product1 := Product{
                ID:    1,
                Name:  "Laptop",
                Price: 1200.50,
        }

        fmt.Println("\nProduct 1:")
        fmt.Println("ID:", product1.ID)
        fmt.Println("Name:", product1.Name)
        fmt.Println("Price:", product1.Price)

        anonymousPerson := struct {
                FirstName string
                LastName  string
                Age       int
        }{
                FirstName: "Jane",
                LastName:  "Smith",
                Age:       25,
        }

        fmt.Println("\nAnonymous Person:")
        fmt.Println("First Name:", anonymousPerson.FirstName)
        fmt.Println("Last Name:", anonymousPerson.LastName)
        fmt.Println("Age:", anonymousPerson.Age)

        personPtr := &person1
        personPtr.Age = 31

        fmt.Println("\nPerson 1 Age after pointer modification:", person1.Age)

        person2 := Person{"Bob","Johnson", 40, Address{"456 Oak Ave", "Otherville", "67890"}}
        fmt.Println("\nPerson 2:")
        fmt.Println("First Name:", person2.FirstName)
        fmt.Println("Last Name:", person2.LastName)
        fmt.Println("Age:", person2.Age)
        fmt.Println("Street:", person2.Address.Street)
        fmt.Println("City:", person2.Address.City)
        fmt.Println("Zip Code:", person2.Address.ZipCode)
}