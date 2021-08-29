package main

import (
	"log"
	"net/http"

	"spaced.com/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/items/id/{id}", handler.ItemByIdHandler).Methods("GET")
	r.HandleFunc("/items/tag/{tag}", handler.ItemsByTagHandler).Methods("GET")

	r.HandleFunc("/items/", handler.AddItemHandler).Methods("POST")

	r.HandleFunc("/items/{id}", handler.DeleteItemHandler).Methods("DELETE")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
