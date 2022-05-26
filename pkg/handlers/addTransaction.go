package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"example.com/RESTapi/pkg/mocks"
	"example.com/RESTapi/pkg/models"
)

func AddTransaction(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var transaction models.Transaction
	json.Unmarshal(body, &transaction)

	transaction.Id = rand.Intn(100)
	mocks.Transactions = append(mocks.Transactions, transaction)

	w.Header().Add("Content-type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Added")
}
