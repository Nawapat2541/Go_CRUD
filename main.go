package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Book struct {
	ID     int     `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

// all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var x, _ = strconv.Atoi(params["id"]);
	for _, item := range books {
		if x == item.ID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// create
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

// update
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var x, _ = strconv.Atoi(params["id"]);
	for index, item := range books {
		if x == item.ID {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = item.ID
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// delete
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var x, _ = strconv.Atoi(params["id"]);
	for index, item := range books {
		if x == item.ID {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: 1, Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "Somchai", Lastname: "Jaidee"}})
	books = append(books, Book{ID: 2, Isbn: "448744", Title: "Book Two", Author: &Author{Firstname: "Yupaporn", Lastname: "Tunrak"}})
	books = append(books, Book{ID: 3, Isbn: "448745", Title: "Book Three", Author: &Author{Firstname: "Somsak", Lastname: "Srisanti"}})
	books = append(books, Book{ID: 4, Isbn: "448746", Title: "Book Four", Author: &Author{Firstname: "Treeanart", Lastname: "Kampanat"}})
	books = append(books, Book{ID: 5, Isbn: "448747", Title: "Book Five", Author: &Author{Firstname: "Darenee", Lastname: "Champachit"}})
	books = append(books, Book{ID: 6, Isbn: "448748", Title: "Book Six", Author: &Author{Firstname: "Sitti", Lastname: "Satitvithi"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
