package handler

import (
	"fmt"
	"net/http"
)

func ItemByIdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ItmByIdHandler")
}

func ItemsByTagHandler(w http.ResponseWriter, r *http.Request) {
}

func AddItemHandler(w http.ResponseWriter, r *http.Request) {
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
}
