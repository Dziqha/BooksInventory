package routers

import (
    "BooksInventory/src/controllers"
    "github.com/gofiber/fiber/v2"
)

func NewRouterBooks(router fiber.Router, booksController *controllers.BooksController) {

    booksgruop := router.Group("/api/v1")
    booksgruop.Post("/books", booksController.Add)
    booksgruop.Put("/books/:id", booksController.Update)
    booksgruop.Get("/books", booksController.FindAll)
    booksgruop.Get("/books/:id", booksController.FindById)
    booksgruop.Delete("/books/:id", booksController.Delete)
}
