package main

import (
	"log"
	"net/http"

	"example.com/RESTapi/pkg/db"
	"example.com/RESTapi/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/products/{id}", h.GetAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/cart/products", h.AddProduct).Methods(http.MethodPost)
	router.HandleFunc("/cart/products", h.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/cart/products/{id}", h.DeleteProduct).Methods(http.MethodDelete)
	router.HandleFunc("/transactions", h.AddTransaction).Methods(http.MethodPost)
	router.HandleFunc("/users", h.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", h.GetUser).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
