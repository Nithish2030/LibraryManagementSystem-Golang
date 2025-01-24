package library

import (
	"errors"
	"fmt"
)

type Library struct {
	Books []Book
}

func (lib *Library) AddBook(book Book) error {
	for _, b := range lib.Books {
		if b.ISBN == book.ISBN {
			return errors.New("duplicate ISBN")
		}
	}
	lib.Books = append(lib.Books, book)
	return nil
}

func (lib *Library) RemoveBook(isbn string) error {
	for i, b := range lib.Books {
		if b.ISBN == isbn {
			lib.Books = append(lib.Books[:i], lib.Books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}

func (lib *Library) SearchBookByTitle(title string) *Book {
	for _, b := range lib.Books {
		if b.Title == title {
			return &b
		}
	}
	return nil
}

func (lib *Library) ListBooks() {
	for _, b := range lib.Books {
		fmt.Printf("Title: %s, Author: %s, ISBN: %s\n", b.Title, b.Author, b.ISBN)
		if b.IsEbook {
			fmt.Println("This book is an eBook")
		}
	}
}
