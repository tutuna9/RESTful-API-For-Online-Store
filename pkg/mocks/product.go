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

var Carts = []models.Cart{
	{
		Id:   1,
		Name: "Sandals",
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
	{
		Id:   2,
		Name: "Budo",
	},
}
