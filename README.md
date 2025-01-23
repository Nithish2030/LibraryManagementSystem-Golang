# LibraryManagementSystem-Golang


Below is the complete code for the Library Management System based on your assignment, split into appropriate files and directories. I’ll also provide the filenames and how the structure should look like.

### Project Directory Structure:
```
/LibraryManagementSystem
│
├── main.go                // Main application logic and entry point
├── book.go                // Definitions and methods related to Book struct
├── ebook.go               // Definitions and methods related to EBook struct
├── library.go             // Definitions and methods related to Library struct
└── README.md              // Documentation file with instructions
```

### 1. **`main.go`** — Main application logic

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var library LibraryEnhanced
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Display menu options
		fmt.Println("Library Management System")
		fmt.Println("1. Add Book/EBook")
		fmt.Println("2. Remove Book/EBook by ISBN")
		fmt.Println("3. Search for Books by Title")
		fmt.Println("4. List all Books/EBooks")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		// Read user choice
		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			// Add a Book/EBook
			fmt.Print("Enter book type (book/ebook): ")
			scanner.Scan()
			bookType := scanner.Text()

			if bookType == "book" {
				// Add a regular book
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

			} else if bookType == "ebook" {
				// Add an eBook
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

		case "2":
			// Remove a Book/EBook by ISBN
			fmt.Print("Enter ISBN to remove: ")
			scanner.Scan()
			isbn := scanner.Text()
			err := library.RemoveBook(isbn)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book/Ebook removed successfully!")
			}

		case "3":
			// Search for books by title
			fmt.Print("Enter title to search: ")
			scanner.Scan()
			title := scanner.Text()
			books := library.SearchBookByTitle(title)
			for _, book := range books {
				fmt.Println(book.DisplayDetails())
			}

		case "4":
			// List all Books/EBooks
			library.ListBooks()

		case "5":
			// Exit the program
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
```

### 2. **`book.go`** — Book struct and methods

```go
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
```

### 3. **`ebook.go`** — EBook struct and methods

```go
package main

import "fmt"

// EBook struct, embedding Book
type EBook struct {
	Book
	FileSize string // Size of the file in MB
}

// Override the DisplayDetails method for EBook
func (e EBook) DisplayDetails() string {
	return fmt.Sprintf("Title: %s\nAuthor: %s\nISBN: %s\nAvailable: %t\nFile Size: %s MB\n",
		e.Title, e.Author, e.ISBN, e.Available, e.FileSize)
}
```

### 4. **`library.go`** — Library struct and methods

```go
package main

import (
	"fmt"
	"strings"
)

// BookInterface defines the common methods for Book and EBook
type BookInterface interface {
	DisplayDetails() string
}

// LibraryEnhanced struct to manage a collection of books (both Book and EBook)
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
```

### 5. **`README.md`** — Documentation file with instructions

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

Title: Advanced Go Programming
Author: Jane Smith
ISBN: 987654321
Available: true
File Size: 5 MB
```

### Remove a Book by ISBN:
```
Enter ISBN to remove: 123456789
Book/Ebook removed successfully!
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

---

### Instructions for Setup and Running:

1. **Install Go**: Make sure you have Go installed on your system. If you haven't installed Go, follow the instructions here: [Installing Go](https://golang.org/doc/install).

2. **Clone/Download the Project**: Save the files to your computer, and make sure the directory structure looks like this:

```
LibraryManagementSystem/
├── main.go
├── book.go
├── ebook.go
├── library.go
└── README.md
```

3. **Run the Program**:
   - Open a terminal and navigate to the project directory:
     ```bash
     cd path/to/LibraryManagementSystem
     ```
   - Run the program:
     ```bash
     go run main.go
     ```

4. **Interact with the Program**: Follow the on-screen instructions to add books, search, remove, and list books in the library.

---

With this structure

, you should be able to demonstrate all required features of the Library Management System as per the assignment requirements. Let me know if you need further help!
