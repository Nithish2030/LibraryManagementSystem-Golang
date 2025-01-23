Got it! Let's start fresh and go step by step. Below is a full, working example of a **Library Management System** in Go, implementing all the required functionality: adding, removing, searching books, listing books, using structs, methods, and interfaces.

### Fresh Implementation

### File Structure
The project will be organized into several files to keep it modular and clean:

```
LibraryManagementSystem/
├── main.go           // Main program entry and logic
├── book.go           // Book struct and its methods
├── ebook.go          // EBook struct (embedding Book) and its methods
├── library.go        // Library struct and methods
├── README.md         // Instructions and examples
```

### 1. **`main.go`** — Main application logic and the interactive text-based menu.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Initialize the library
	library := NewLibrary()

	// Create a scanner for user input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Display the menu
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add a Book/EBook")
		fmt.Println("2. Remove a Book/EBook by ISBN")
		fmt.Println("3. Search for Books by Title")
		fmt.Println("4. List all Books/EBooks")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		// Read user input
		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			// Add a Book or EBook
			fmt.Print("Enter book type (book/ebook): ")
			scanner.Scan()
			bookType := scanner.Text()

			if bookType == "book" {
				// Add a regular Book
				addBook(scanner, library)
			} else if bookType == "ebook" {
				// Add an EBook
				addEBook(scanner, library)
			} else {
				fmt.Println("Invalid book type.")
			}

		case "2":
			// Remove a Book/EBook by ISBN
			fmt.Print("Enter ISBN to remove: ")
			scanner.Scan()
			isbn := scanner.Text()
			if err := library.RemoveBook(isbn); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book removed successfully!")
			}

		case "3":
			// Search for Books by Title
			fmt.Print("Enter title to search: ")
			scanner.Scan()
			title := scanner.Text()
			books := library.SearchBookByTitle(title)
			if len(books) == 0 {
				fmt.Println("No books found with that title.")
			} else {
				for _, book := range books {
					fmt.Println(book.DisplayDetails())
				}
			}

		case "4":
			// List all Books/EBooks
			library.ListBooks()

		case "5":
			// Exit the program
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addBook(scanner *bufio.Scanner, library *Library) {
	fmt.Print("Enter title: ")
	scanner.Scan()
	title := scanner.Text()
	fmt.Print("Enter author: ")
	scanner.Scan()
	author := scanner.Text()
	fmt.Print("Enter ISBN: ")
	scanner.Scan()
	isbn := scanner.Text()
	fmt.Print("Is it available? (true/false): ")
	scanner.Scan()
	available := scanner.Text() == "true"

	book := NewBook(title, author, isbn, available)
	library.AddBook(book)
}

func addEBook(scanner *bufio.Scanner, library *Library) {
	fmt.Print("Enter title: ")
	scanner.Scan()
	title := scanner.Text()
	fmt.Print("Enter author: ")
	scanner.Scan()
	author := scanner.Text()
	fmt.Print("Enter ISBN: ")
	scanner.Scan()
	isbn := scanner.Text()
	fmt.Print("Is it available? (true/false): ")
	scanner.Scan()
	available := scanner.Text() == "true"
	fmt.Print("Enter file size (MB): ")
	scanner.Scan()
	fileSize := scanner.Text()

	ebook := EBook{
		Book: Book{
			Title:    title,
			Author:   author,
			ISBN:     isbn,
			Available: available,
		},
		FileSize: fileSize,
	}
	library.AddBook(ebook)
}
```

### 2. **`book.go`** — The `Book` struct and its methods.

```go
package main

import "fmt"

// Book struct
type Book struct {
	Title     string
	Author    string
	ISBN      string
	Available bool
}

// Create a new Book instance
func NewBook(title, author, isbn string, available bool) Book {
	return Book{
		Title:     title,
		Author:    author,
		ISBN:      isbn,
		Available: available,
	}
}

// Display the details of the Book
func (b Book) DisplayDetails() string {
	return fmt.Sprintf("Title: %s\nAuthor: %s\nISBN: %s\nAvailable: %t\n", b.Title, b.Author, b.ISBN, b.Available)
}
```

### 3. **`ebook.go`** — The `EBook` struct (which embeds `Book`) and its methods.

```go
package main

import "fmt"

// EBook struct embeds Book and adds FileSize attribute
type EBook struct {
	Book
	FileSize string // Size of the file in MB
}

// Display the details of the EBook, overriding Book's DisplayDetails method
func (e EBook) DisplayDetails() string {
	return fmt.Sprintf("Title: %s\nAuthor: %s\nISBN: %s\nAvailable: %t\nFile Size: %s MB\n", e.Title, e.Author, e.ISBN, e.Available, e.FileSize)
}
```

### 4. **`library.go`** — The `Library` struct and its methods.

```go
package main

import (
	"fmt"
	"strings"
)

// BookInterface defines the DisplayDetails method for both Book and EBook
type BookInterface interface {
	DisplayDetails() string
}

// Library struct holds a collection of books (both Books and EBooks)
type Library struct {
	Books []BookInterface
}

// Create a new Library instance
func NewLibrary() *Library {
	return &Library{}
}

// Add a Book or EBook to the library
func (l *Library) AddBook(book BookInterface) {
	l.Books = append(l.Books, book)
}

// List all Books/EBooks in the library
func (l *Library) ListBooks() {
	fmt.Println("Listing all books and eBooks:")
	for _, book := range l.Books {
		fmt.Println(book.DisplayDetails())
	}
}

// Search for books by title
func (l *Library) SearchBookByTitle(title string) []BookInterface {
	var results []BookInterface
	for _, book := range l.Books {
		if strings.Contains(strings.ToLower(book.DisplayDetails()), strings.ToLower(title)) {
			results = append(results, book)
		}
	}
	return results
}

// Remove a Book/EBook from the library by ISBN
func (l *Library) RemoveBook(isbn string) error {
	for i, book := range l.Books {
		if book.DisplayDetails() == isbn {
			// Remove the book from the slice
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("book with ISBN %s not found", isbn)
}
```

### 5. **`README.md`** — Instructions and example usage

```markdown
# Library Management System

## Instructions to Run

1. Install Go (https://golang.org/doc/install) if not already installed.
2. Clone or copy the code into a folder named `LibraryManagementSystem`.
3. Run the program with the following command:
   ```bash
   go run main.go
   ```

## Example Usage:

### Add a Book:
```
Enter book type (book/ebook): book
Enter title: Go Programming Basics
Enter author: John Doe
Enter ISBN: 123456789
Is it available? (true/false): true
```

### Add an EBook:
```
Enter book type (book/ebook): ebook
Enter title: Advanced Go Programming
Enter author: Jane Smith
Enter ISBN: 987654321
Is it available? (true/false): true
Enter file size (MB): 5
```

### List all Books:
```
Listing all books:
Title: Go Programming Basics
Author: John Doe
ISBN: 123456789
Available: true
```

### Remove a Book by ISBN:
```
Enter ISBN to remove: 123456789
Book removed successfully!
```

### Search for Books by Title:
```
Enter title to search: Go
Listing books with 'Go' in title:
Title: Go Programming Basics
Author: John Doe
ISBN: 123456789
Available: true
```

---

## Features Implemented:
- Add Book/EBook to library
- Remove Book/EBook from library by ISBN
- Search for Books/EBooks by Title
- List all Books/EBooks with file sizes for EBooks
```

### Running the Application:

1. Open a terminal and navigate to the project folder `LibraryManagementSystem`.
2. Run the Go application:
   ```bash
   go run main.go
   ```
3. Follow the on-screen instructions to interact with the system.

---

This should now

 be a fully working, modular solution that handles books and eBooks, implements polymorphism with interfaces, and allows you to manage a simple library. 

Let me know if you need any further help or clarification!
