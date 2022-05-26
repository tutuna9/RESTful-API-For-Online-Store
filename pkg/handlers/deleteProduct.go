package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"example.com/RESTapi/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var cart models.Cart
	
	if result := h.DB.First(&cart, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&cart)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}