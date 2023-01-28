package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "you've request the book: %s on page %s\n", title, page)
	})

	CreateBook := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "you've create the book: %s\n", title)
	}
	r.HandleFunc("/books/{title}/create", CreateBook).Methods("GET")

	AllBooks := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is all books")
	}

	GetBook := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "you've get the book: %s\n", title)
	}

	bookrouter := r.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/", AllBooks)
	bookrouter.HandleFunc("/{title}", GetBook)

	http.ListenAndServe(":80", r)
}
