package main

import "fmt"

// EBook struct, embedding Book
type EBook struct {
	Book
	FileSize int // Size of the file in MB
}

// Override the DisplayDetails method for EBook
func (e EBook) DisplayDetails() string {
	return fmt.Sprintf("Title: %s\nAuthor: %s\nISBN: %s\nAvailable: %t\nFile Size: %d MB\n",
		e.Title, e.Author, e.ISBN, e.Available, e.FileSize)
}
