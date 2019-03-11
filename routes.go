package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc("/authors", AuthorIndex)
	router.HandleFunc("/authors/{authorID}", AuthorShow)

	return router
}
