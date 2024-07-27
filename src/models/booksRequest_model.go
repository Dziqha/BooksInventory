package models

type BooksRequestCreate struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	ISBN   string `json:"isbn" validate:"required"`
	Amoutnt int `json:"amount" validate:"required"`
}

type BooksRequestUpdate struct {
	Id     int    `json:"id" validate:"required"`
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	ISBN   string `json:"isbn" validate:"required"`
	Amoutnt int `json:"amount" validate:"required"`
}
