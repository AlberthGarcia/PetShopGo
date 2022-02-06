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
	mux.HandleFunc("/api/getCategory/{id:[0-9]+}", handlers.GetCategoryById).Methods("GET")
	mux.HandleFunc("/api/createCategory/", handlers.CreateCategory).Methods("POST")
	mux.HandleFunc("/api/deleteCategory/{id:[0-9]+}", handlers.DeleteCategory).Methods("Delete")
	mux.HandleFunc("/api/updateCategory/{id:[0-9]+}", handlers.UpdateCategory).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", mux))
}
