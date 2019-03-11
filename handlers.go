package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Empty index page")
}

func AuthorIndex(w http.ResponseWriter, r *http.Request) {
	authors := Authors{
		Author{Name: "Jack London"},
		Author{Name: "Mark Twain"},
		Author{Name: "Franz Kafka"},
	}

	if err := json.NewEncoder(w).Encode(authors); err != nil {
		panic(err)
	}
}

func AuthorShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorID := vars["authorID"]
	fmt.Fprintf(w, "Show Author with ID %s", authorID)
}
