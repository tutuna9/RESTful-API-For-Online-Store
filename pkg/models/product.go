package models

type Product struct {
	CategoryId int    `json:"categoryid"`
	Id         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
}

type Cart struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Transaction struct {
	Id int `json:"id" gorm:"primaryKey"`
}

type User struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
