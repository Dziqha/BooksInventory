package main

import (
	"BooksInventory/app"
	"BooksInventory/app/middleware"
	"BooksInventory/src/controllers"
	"BooksInventory/src/routers"
	"BooksInventory/src/services"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	ap := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 10,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Prefork:      true,
	})

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",

	})

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	apikey := os.Getenv("API_KEY")

	if apikey == "" {
		log.Fatal("API_KEY is not set in the environment")
	}
	
	db := app.Database() 

	validate := validator.New()

	booksService := services.NewBooksService(db)

	booksController := controllers.NewBooksController(booksService, validate, rdb)

	ap.Use(logger.New())
	ap.Use(recover.New())
	ap.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	ap.Use(middleware.ApiKeyMiddleware(apikey))

	routers.NewRouterBooks(ap, booksController)

	err := ap.Listen("localhost:3000")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
