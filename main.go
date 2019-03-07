package main

import (
	"encoding/json"
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
	authors := Authors{
		Author{Name: "Jack London"},
		Author{Name: "Mark Twain"},
		Author{Name: "Franz Kafka"},
	}

	json.NewEncoder(w).Encode(authors)
}

func authorShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorID := vars["authorID"]
	fmt.Fprintf(w, "Show Author with ID %s", authorID)
}
