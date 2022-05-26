package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/RESTapi/pkg/models"
)

func (h handler) AddTransaction(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var transaction models.Transaction
	json.Unmarshal(body, &transaction)

	if result := h.DB.Create(&transaction); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Added")
}
