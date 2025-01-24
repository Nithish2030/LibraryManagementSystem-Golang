package main

import (
	"LMS_GO/src/library" // Module-based import path
)

func main() {
	libraryInstance := library.Library{}
	library.Menu(&libraryInstance)
}
