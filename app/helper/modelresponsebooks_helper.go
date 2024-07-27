package helper

import "BooksInventory/src/models"

func ToResponseBooks(books models.Books) models.BooksResponse {
	return models.BooksResponse{
		Id: books.Id,
		Title: books.Title,
		Author: books.Author,
		Amount: books.Amount,
	}
}
