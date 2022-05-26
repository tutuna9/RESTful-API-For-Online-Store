package handlers

import (
	"encoding/json"
	"net/http"

	"example.com/RESTapi/pkg/mocks"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Products)
}
