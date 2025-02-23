package main

import (
	"fmt"
)

func main() {
	InitDB() // Initialize the database

	// Example usage (replace with your desired application logic)
	book1 := Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams"}
	createdBook, err := CreateBook(book1)
	if err != nil {
		fmt.Println("Error creating book:", err)
		return
	}
	fmt.Println("Created book:", createdBook)

	// Update the created book
	createdBook.Author = "Douglas Adams"
	updatedBook, err := UpdateBook(createdBook.ID, createdBook)
	if err != nil {
		fmt.Println("Error updating book:", err)
		return
	}
	fmt.Println("Updated book:", updatedBook)

	books, err := GetBooks()
	if err != nil {
		fmt.Println("Error getting books:", err)
		return
	}
	fmt.Println("All books:", books)

	// Get a book by ID
	book, err := GetBook(1)
	if err != nil {
		fmt.Println("Error getting book:", err)
		return
	}
	fmt.Println("Book with ID 1:", book)


	err = DeleteBook(createdBook.ID)
	if err != nil {
		fmt.Println("Error deleting book:", err)
		return
	}
	fmt.Println("Book deleted successfully.")
}
