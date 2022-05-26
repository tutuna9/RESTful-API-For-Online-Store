package handlers

import (
	"encoding/json"
	"net/http"

	"example.com/RESTapi/pkg/mocks"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Conent-Tpe", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Carts)
}
	