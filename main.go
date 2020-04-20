package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	router := chi.NewRouter()
	http.ListenAndServe(":8000", nil)
}
