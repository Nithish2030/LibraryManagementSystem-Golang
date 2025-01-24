package library

type Book struct {
	Title   string
	Author  string
	ISBN    string
	IsEbook bool
}

type EBook struct {
	Book
	FileSize int
}

type BookOperation interface {
	AddBook(book Book) error
	RemoveBook(isbn string) error
	SearchBookByTitle(title string) *Book
	ListBooks()
}
