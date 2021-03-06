package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"math/rand"
)

//Book Struct (Model)
type Book struct{
	ID			string `json:"id"`
	Isbn		string `json:"isbn"`
	Title	  string `json:"title"`
	Author *Author `json:"author"`
}

//Author Struct
type Author struct{
	Firstname string `json:"firstname"`
	Lastname	string `json:"lastname"`
}

//inits books var as a slice book struct
var books []Book

//get all books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
//get single book
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	//get params
	params := mux.Vars(r)
	//loop through books and find with it
	for _, item := range books{
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}
//create new book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) //mock id not safer
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
//update book
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"]{
		books = append(books[:index], books[index+1:]...)
		var book Book
		_ = json.NewDecoder(r.Body).Decode(&book)
		book.ID = params["id"]
		books = append(books, book)
		json.NewEncoder(w).Encode(book)
		return
		}
	}
	json.NewEncoder(w).Encode(books)
}
//delete book
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"]{
		books = append(books[:index], books[index+1:]...)
		break
		}
	}
	json.NewEncoder(w).Encode(books)

}

func main() {
	//init router
	r := mux.NewRouter()

	//mock data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "Hafiz", Lastname: "Syap"}})

	books = append(books, Book{ID: "2", Isbn: "847653", Title: "Book Two", Author: &Author{Firstname: "Yusuf", Lastname: "Syap"}})

	//route handlers/ endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

	
}