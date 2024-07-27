package models


type BooksResponse struct {
	Id int		`json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	ISBN string `json:"isbn"`
	Amount int `json:"amount"`
}
