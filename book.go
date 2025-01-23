package main

import "fmt"

// Book struct
type Book struct {
	Title    string
	Author   string
	ISBN     string
	Available bool
}

// Method to initialize a new book
func NewBook(title, author, isbn string, available bool) Book {
	return Book{
		Title:    title,
		Author:   author,
		ISBN:     isbn,
		Available: available,
	}
}

// Method to display book details
func (b Book) DisplayDetails() string {
	return fmt.Sprintf("Title: %s\nAuthor: %s\nISBN: %s\nAvailable: %t\n", b.Title, b.Author, b.ISBN, b.Available)
}
