package main

import (
	"fmt"
)

type Book struct {
	BookID int
	Title  string
	Author string
	Price  float64
}

func main() {
	var n int
	fmt.Print("Enter the number of books: ")
	fmt.Scan(&n)

	books := make([]Book, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter details for book %d:\n", i+1)
		fmt.Print("Book ID: ")
		fmt.Scan(&books[i].BookID)
		fmt.Print("Title: ")
		fmt.Scan(&books[i].Title)
		fmt.Print("Author: ")
		fmt.Scan(&books[i].Author)
		fmt.Print("Price: ")
		fmt.Scan(&books[i].Price)
	}

	fmt.Println("\nBook Details:")
	for _, book := range books {
		fmt.Printf("Book ID: %d, Title: %s, Author: %s, Price: %.2f\n", book.BookID, book.Title, book.Author, book.Price)
	}
}
