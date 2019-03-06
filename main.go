package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Coming soon...")
	})
	
	log.Fatal(http.ListenAndServe(":8000", nil))
}
