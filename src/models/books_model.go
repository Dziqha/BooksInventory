package models

type Books struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Amount int `json:"amount"`
}