package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/RESTapi/pkg/models"
)

func (h handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	var carts []models.Cart

	if result := h.DB.Find(&carts); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Conent-Tpe", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(carts)
}
