package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/RESTapi/pkg/models"
)

func (h handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var cart models.Cart
	json.Unmarshal(body, &cart)

	if result := h.DB.Create(&cart); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-type", "alication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Added")
}
