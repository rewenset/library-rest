package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Logger(Index))
	router.HandleFunc("/authors", Logger(AuthorIndex))
	router.HandleFunc("/authors/{authorID}", Logger(AuthorShow))

	return router
}
