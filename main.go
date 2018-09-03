package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book Struck (Model)
type Book struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
}

//Init books var as slice book struct
var books []Book

//Controller Get All Books
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)

}

//controller Get Book By ID
func getBooksByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r) //Get params

	//Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) //mock id =not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

//Router
func main() {
	//Init Router
	r := mux.NewRouter()

	//Mock Data todo - Implement DB
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Pendidikan Agama Islam"})
	books = append(books, Book{ID: "2", Isbn: "448744", Title: "Pendidikan Agama Kristen"})

	//Route Handlers/ Endpoint
	r.HandleFunc("/api/books", getAllBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBooksByID).Methods("GET")
	r.HandleFunc("/api/books/create", CreateBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
