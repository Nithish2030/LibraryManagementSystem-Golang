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
