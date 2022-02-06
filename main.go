package main

import (
	"log"
	"net/http"
	"package/handlers"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/api/getCategories/", handlers.GetCategories).Methods("GET")
	mux.HandleFunc("/api/getCategory/{id:[0-9]+}", nil).Methods("GET")
	mux.HandleFunc("/api/createCategory/", nil).Methods("POST")
	mux.HandleFunc("/api/deleteCategory/{id:[0-9]+}", nil).Methods("Delete")
	mux.HandleFunc("/api/updateCategory/{id:[0-9]+}", nil).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", mux))
}
