package db

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var (
	mutex    sync.Mutex
	filePath = "books.json"
)

func readBooks() ([]Book, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var books []Book
	if err := json.Unmarshal(file, &books); err != nil {
		return nil, err
	}
	return books, nil
}

func writeBooks(books []Book) error {
	data, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func AddBookToDB(bookID, status string) {
	mutex.Lock()
	defer mutex.Unlock()

	books, err := readBooks()
	if err != nil {
		fmt.Println("Error reading books:", err)
		return
	}

	book := Book{
		ID:     bookID,
		Status: status,
	}
	books = append(books, book)

	if err := writeBooks(books); err != nil {
		fmt.Println("Error writing books:", err)
		return
	}

	fmt.Println("Book added successfully!")
}
