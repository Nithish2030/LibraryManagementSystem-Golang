package main

import (
	"fmt"
	"strings"
)

// Library struct to manage a collection of books (both Book and EBook)
type LibraryEnhanced struct {
	Books []BookInterface
}

// Add a Book/EBook to the library
func (l *LibraryEnhanced) AddBook(book BookInterface) {
	l.Books = append(l.Books, book)
}

// List all Books/EBooks
func (l *LibraryEnhanced) ListBooks() {
	for _, book := range l.Books {
		fmt.Println(book.DisplayDetails())
	}
}

// Method to search for books by title
func (l *LibraryEnhanced) SearchBookByTitle(title string) []BookInterface {
	var results []BookInterface
	for _, book := range l.Books {
		if strings.Contains(strings.ToLower(book.DisplayDetails()), strings.ToLower(title)) {
			results = append(results, book)
		}
	}
	return results
}

// Remove a Book/EBook by ISBN
func (l *LibraryEnhanced) RemoveBook(isbn string) error {
	for i, book := range l.Books {
		if book.DisplayDetails() == isbn {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("book with ISBN %s not found", isbn)
}
