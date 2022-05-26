package mocks

import "example.com/RESTapi/pkg/models"

var Products = []models.Product{
	{
		CategoryId: 1,
		Id:         1,
		Name:       "Handphone",
	},
	{
		CategoryId: 2,
		Id:         1,
		Name:       "Bag",
	},
}

var Transactions = []models.Transaction{
	{
		Id: 1,
	},
}

var Users = []models.User{
	{
		Id:   1,
		Name: "Budi",
	},
}
