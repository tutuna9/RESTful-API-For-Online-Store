package main

import (
	"log"
	"net/http"

	"example.com/RESTapi/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/products/{id}", handlers.GetAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/cart/products", handlers.AddProduct).Methods(http.MethodPost)
	router.HandleFunc("/cart/products", handlers.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/cart/products/{id}", handlers.DeleteProduct).Methods(http.MethodDelete)
	router.HandleFunc("/transactions", handlers.AddTransaction).Methods(http.MethodPost)
	router.HandleFunc("/users", handlers.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
