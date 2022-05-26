package models

type Product struct {
	CategoryId int    `json:"categoryid"`
	Id         int    `json:"id"`
	Name       string `json:"name"`
}

type Cart struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Transaction struct {
	Id int `json:"id"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
