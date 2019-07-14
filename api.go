package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func returnAllBooksHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint returnAllBooks")
	json.NewEncoder(w).Encode(Books)
}

func createBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Create Book")

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var book Book
	err = json.Unmarshal(requestBody, &book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = createBook(&book)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func getBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Println("Endpoint Get Book ", key)

	id, err := strconv.Atoi(key)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid key.", http.StatusBadRequest)
		return
	}

	book := getBook(key)
	if book == (Book{}) {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func updateBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Println("Endpoint Update Book ", key)

	id, err := strconv.Atoi(key)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid key.", http.StatusBadRequest)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var book Book
	err = json.Unmarshal(requestBody, &book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validateBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = updateBook(key, &book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Println("Endpoint Delete Book ", key)

	id, err := strconv.Atoi(key)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid key.", http.StatusBadRequest)
		return
	}

	err = deleteBook(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Book API HomePage!")
	fmt.Fprintln(w, "POST http://localhost:8084/book    Creates a new book")
	fmt.Fprintln(w, "GET http://localhost:8084/book/{id}    Retrieves book information")
	fmt.Fprintln(w, "PUT http://localhost:8084/book/{id}    Updates book information")
	fmt.Fprintln(w, "DELETE http://localhost:8084/book/{id}    Deletes a book")
	fmt.Println("Endpoint HomePage")
}

func handleRequests() {

	// a new instance of a mux router
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/", homePage)
	muxRouter.HandleFunc("/books", returnAllBooksHandler)

	muxRouter.HandleFunc("/book", createBookHandler).Methods("POST")
	muxRouter.HandleFunc("/book/{id}", getBookHandler).Methods("GET")
	muxRouter.HandleFunc("/book/{id}", updateBookHandler).Methods("PUT")
	muxRouter.HandleFunc("/book/{id}", deleteBookHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8084", muxRouter))
}
