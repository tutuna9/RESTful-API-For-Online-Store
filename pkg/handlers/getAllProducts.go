package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"example.com/RESTapi/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
