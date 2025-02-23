package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite3 driver
)

var db *sql.DB

// Initialize the database
func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "books.db") // Create or open books.db
	if err != nil {
		log.Fatal(err)
	}

	// Create the books table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			author TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database initialized.")
}

// Get all books from the database
func GetBooks() ([]Book, error) {
	rows, err := db.Query("SELECT id, title, author FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

// Get a single book by ID
func GetBook(id int) (Book, error) {
	var book Book
	err := db.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		return Book{}, err // Return an empty Book struct and the error
	}
	return book, nil
}

// Create a new book in the database
func CreateBook(book Book) (Book, error) {
    result, err := db.Exec("INSERT INTO books (title, author) VALUES (?, ?)", book.Title, book.Author)
    if err != nil {
        return Book{}, err
    }

    // Retrieve the last inserted ID (the newly created book's ID)
    lastInsertID, err := result.LastInsertId()
    if err != nil {
        return Book{}, err
    }

    // Update the book's ID with the database-assigned ID
    book.ID = int(lastInsertID)
    return book, nil
}


// Update an existing book
func UpdateBook(id int, book Book) (Book, error) {
	_, err := db.Exec("UPDATE books SET title = ?, author = ? WHERE id = ?", book.Title, book.Author, id)
	if err != nil {
		return Book{}, err
	}
	book.ID = id // Keep the original ID
	return book, nil
}

// Delete a book
func DeleteBook(id int) error {
	_, err := db.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}
