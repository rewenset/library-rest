package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/authors", authorIndex)
	router.HandleFunc("/authors/{authorID}", authorShow)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Empty index page")
}

func authorIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Show all authors")
}

func authorShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorID := vars["authorID"]
	fmt.Fprintf(w, "Show Author with ID %s", authorID)
}
