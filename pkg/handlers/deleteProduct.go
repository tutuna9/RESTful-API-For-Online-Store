package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/RESTapi/pkg/mocks"
	"github.com/gorilla/mux"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, cart := range mocks.Carts {
		if cart.Id == id {
			mocks.Carts = append(mocks.Carts[:index], mocks.Carts[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}
