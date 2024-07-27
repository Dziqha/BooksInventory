package services

import (
	"BooksInventory/src/models"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type BooksService struct {
	DB *sql.DB
}

func NewBooksService(db *sql.DB) BooksService {
	return BooksService{DB: db}
}

func (service *BooksService) Add(ctx context.Context, req models.BooksRequestCreate) models.BooksResponse {
	var res models.BooksResponse
	query := `INSERT INTO books (title, author, isbn, amount) VALUES (?, ?, ?, ?)`
	result, err := service.DB.ExecContext(ctx, query, req.Title, req.Author, req.Amoutnt, req.ISBN)
	if err != nil {
		fmt.Println(err)
		return res
	}
	id, _ := result.LastInsertId()
	res.Id = int(id)
	res.Title = req.Title
	res.Author = req.Author
	res.Amount = req.Amoutnt
	res.ISBN = req.ISBN
	return res
}

func (service *BooksService) Update(ctx context.Context, req models.BooksRequestUpdate) models.BooksResponse {
	var res models.BooksResponse

	querySelect := `SELECT title, author, isbn, amount FROM books WHERE id = ?`
	err := service.DB.QueryRowContext(ctx, querySelect, req.Id).Scan(&res.Title, &res.Author, &res.ISBN, &res.Amount)
	if err != nil {
		fmt.Println(err)
		return res
	}

	if req.Title == "" {
		req.Title = res.Title
	}
	if req.Author == "" {
		req.Author = res.Author
	}
	if req.ISBN == "" {
		req.ISBN = res.ISBN
	}
	if req.Amoutnt == 0 {
		req.Amoutnt = res.Amount
	}

	queryUpdate := `UPDATE books SET title = ?, author = ?, isbn = ?, amount = ? WHERE id = ?`
	_, err = service.DB.ExecContext(ctx, queryUpdate, req.Title, req.Author, req.ISBN, req.Amoutnt, req.Id)
	if err != nil {
		fmt.Println(err)
		return res
	}

	res.Id = req.Id
	res.Title = req.Title
	res.Author = req.Author
	res.ISBN = req.ISBN
	res.Amount = req.Amoutnt
	
	return res
}

func (service *BooksService) FindAll(ctx context.Context) []models.BooksResponse {
	query := `SELECT * FROM books`
	rows, err := service.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()

	var books []models.BooksResponse
	for rows.Next() {
		var book models.BooksResponse
		if err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.ISBN, &book.Amount); err != nil {
			fmt.Println(err)
			continue
		}
		books = append(books, book)
	}
	return books
}

func (service *BooksService) FindById(ctx context.Context, id int) (models.BooksResponse, error) {
	var book models.BooksResponse
	query := `SELECT id, title, author, isbn, amount FROM books WHERE id = ?`
	row := service.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&book.Id, &book.Title, &book.Author, &book.ISBN, &book.Amount)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, fmt.Errorf("book not found")
		}
		fmt.Println(err)
		return book, fmt.Errorf("error finding book")
	}
	return book, nil
}

func (service *BooksService) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM books WHERE id = ?`
	_, err := service.DB.ExecContext(ctx, query, id)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error deleting book")
	}
	return nil
}
