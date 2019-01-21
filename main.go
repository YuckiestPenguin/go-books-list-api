package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func main() {

	books = append(books, Book{
		ID:     1,
		Title:  "The Life of Vico",
		Author: "Vico",
		Year:   "2019",
	},
		Book{
			ID:     2,
			Title:  "The Life of Roger",
			Author: "Roger",
			Year:   "2019",
		})
	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	//just returns the books list as json for client
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	//grabs the id param from the route
	params := mux.Vars(r)

	//converts the param books/1 from a string to int and assigns to "i"
	i, _ := strconv.Atoi(params["id"])

	//loops through all of the books and checks to see if the route param id matches any ids in the books slice
	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	//init a variable for new book to live it before sending
	var book Book

	//Grab the new book you are creating, grab the json version and turn it go code
	json.NewDecoder(r.Body).Decode(&book)

	//add that new book to books slice
	books = append(books, book)

	//send back the full books list slice as json for client to see
	json.NewEncoder(w).Encode(books)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}

	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	//grabs the id param from the route
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(books)
}
