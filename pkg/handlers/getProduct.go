package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/RESTapi/pkg/mocks"
	"github.com/gorilla/mux"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, product := range mocks.Products {
		if product.Id == id {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(product)
			break
		}
	}
}
