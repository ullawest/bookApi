package main

import (
	"errors"
	"strings"
)

// Book structure
type Book struct {
	ID          string `json:"Id"`
	Title       string `json:"Title"`
	Author      string `json:"Author"`
	Publisher   string `json:"Publisher"`
	PublishDate string `json:"PublishDate"`
	Rating      int    `json:"Rating"`
	Status      string `json:"Status"`
}

var Books []Book

func initializeBooks() {

	Books = []Book{
		Book{ID: "1", Title: "Hello World", Author: "Jane Doe", Publisher: "Publish House", PublishDate: "10/01/2017", Rating: 2, Status: "Published"},
		Book{ID: "2", Title: "Hello City", Author: "John Smith", Publisher: "Publish House", PublishDate: "10/01/2019", Rating: 2, Status: "Under Review"},
	}
}

func validateBook(book *Book) error {
	// check that required fields are not empty
	if len(strings.TrimSpace(book.ID)) == 0 {
		return errors.New("Missing book ID")
	}
	if len(strings.TrimSpace(book.Title)) == 0 {
		return errors.New("Missing book title")
	}
	if len(strings.TrimSpace(book.Author)) == 0 {
		return errors.New("Missing author")
	}
	return nil
}

func createBook(book *Book) error {

	err := validateBook(book)
	if err != nil {
		return err
	}

	Books = append(Books, *book)
	return nil
}

func getBook(key string) Book {

	for _, book := range Books {
		if book.ID == key {
			return book
		}
	}

	return Book{}
}

func updateBook(key string, updatedBook *Book) error {

	for index, book := range Books {
		if book.ID == key {
			Books = append(Books[:index], Books[index+1:]...)
			Books = append(Books, *updatedBook)
			return nil
		}
	}

	return errors.New("Book not found")
}

func deleteBook(key string) error {

	for index, book := range Books {
		if book.ID == key {
			Books = append(Books[:index], Books[index+1:]...)
			return nil
		}
	}

	return errors.New("Book not found")
}
