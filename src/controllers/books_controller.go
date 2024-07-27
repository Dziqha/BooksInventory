package controllers

import (
	"BooksInventory/src/models"
	"BooksInventory/src/services"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // Automatically loads .env
)

type BooksController struct {
    BooksService services.BooksService
    Validate     *validator.Validate
    RedisClient  *redis.Client
}

func NewBooksController(booksService services.BooksService, validate *validator.Validate, redisClient *redis.Client) *BooksController {
    return &BooksController{
        BooksService: booksService,
        Validate:     validate,
        RedisClient:  redisClient,
    }
}

func (controller *BooksController) Add(c *fiber.Ctx) error {
    var req models.BooksRequestCreate
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(models.ResponseCode{
            Code:    fiber.StatusBadRequest,
            Message: "Invalid request body",
            Data:    nil,
        })
    }

    if err := controller.Validate.Struct(req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(models.ResponseCode{
            Code:    fiber.StatusBadRequest,
            Message: "Validation failed: " + err.Error(),
            Data:    nil,
        })
    }

    res := controller.BooksService.Add(c.Context(), req)

    // Clear the cache for books list after adding a new book
    cacheKey := os.Getenv("CACHE_KEY_BOOKS_ALL")
    controller.RedisClient.Del(c.Context(), cacheKey)

    return c.JSON(models.ResponseCode{
        Code:    fiber.StatusOK,
        Message: "Book added successfully",
        Data:    res,
    })
}

func (controller *BooksController) Update(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    var req models.BooksRequestUpdate
    req.Id = id

    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(models.ResponseCode{
            Code:    fiber.StatusBadRequest,
            Message: err.Error(),
            Data:    nil,
        })
    }

    res := controller.BooksService.Update(c.Context(), req)

    // Clear the cache for the specific book after updating
    cacheKey := os.Getenv("CACHE_KEY_BOOKS_PREFIX") + strconv.Itoa(id)
    controller.RedisClient.Del(c.Context(), cacheKey)

    // Clear the cache for books list after updating
    controller.RedisClient.Del(c.Context(), os.Getenv("CACHE_KEY_BOOKS_ALL"))

    return c.JSON(models.ResponseCode{
        Code:    fiber.StatusOK,
        Message: "Book updated successfully",
        Data:    res,
    })
}

func (controller *BooksController) FindAll(c *fiber.Ctx) error {
    startTime := time.Now()

    cacheKey := os.Getenv("CACHE_KEY_BOOKS_ALL")
    val, err := controller.RedisClient.Get(c.Context(), cacheKey).Result()
    if err == nil && val != "" {
        var data []models.BooksResponse
        if err := json.Unmarshal([]byte(val), &data); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseCode{
                Code:    fiber.StatusInternalServerError,
                Message: "Failed to parse cache data",
                Data:    nil,
            })
        }
        c.Set("X-Source", "cache")
        log.Printf("FindAll from cache took %v", time.Since(startTime))
        return c.JSON(models.ResponseCode{
            Code:    fiber.StatusOK,
            Message: "Books retrieved from cache",
            Data:    data,
        })
    }

    books := controller.BooksService.FindAll(c.Context())
    if books == nil {
        return c.Status(fiber.StatusNotFound).JSON(models.ResponseCode{
            Code:    fiber.StatusNotFound,
            Message: "Books not found",
            Data:    nil,
        })
    }

    var responseData []models.BooksResponse
    for _, book := range books {
        responseData = append(responseData, models.BooksResponse{
            Id:     book.Id,
            Title:  book.Title,
            Author: book.Author,
            ISBN:   book.ISBN,
            Amount: book.Amount,
        })
    }

    resJSON, err := json.Marshal(responseData)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseCode{
            Code:    fiber.StatusInternalServerError,
            Message: "Failed to marshal books",
            Data:    nil,
        })
    }

    // Increase cache expiry time to reduce database hits
    err = controller.RedisClient.Set(c.Context(), cacheKey, resJSON, 5 * time.Minute).Err()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseCode{
            Code:    fiber.StatusInternalServerError,
            Message: "Failed to cache books",
            Data:    nil,
        })
    }

    c.Set("X-Source", "database")
    log.Printf("FindAll from database took %v", time.Since(startTime))
    return c.JSON(models.ResponseCode{
        Code:    fiber.StatusOK,
        Message: "Books retrieved successfully",
        Data:    responseData,
    })
}

func (controller *BooksController) FindById(c *fiber.Ctx) error {
    startTime := time.Now()

    id, _ := strconv.Atoi(c.Params("id"))
    cacheKey := os.Getenv("CACHE_KEY_BOOKS_PREFIX") + strconv.Itoa(id)

    val, err := controller.RedisClient.Get(c.Context(), cacheKey).Result()
    if err == nil && val != "" {
        var data models.BooksResponse
        if err := json.Unmarshal([]byte(val), &data); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseCode{
                Code:    fiber.StatusInternalServerError,
                Message: "Failed to parse cache data",
                Data:    nil,
            })
        }
        c.Set("X-Source", "cache")
        log.Printf("FindById from cache took %v", time.Since(startTime))
        return c.JSON(models.ResponseCode{
            Code:    fiber.StatusOK,
            Message: "Book retrieved from cache",
            Data:    data,
        })
    }

    book, err := controller.BooksService.FindById(c.Context(), id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(models.ResponseCode{
            Code:    fiber.StatusNotFound,
            Message: err.Error(),
            Data:    nil,
        })
    }

    responseData := models.BooksResponse{
        Id:     book.Id,
        Title:  book.Title,
        Author: book.Author,
        ISBN:   book.ISBN,
        Amount: book.Amount,
    }

    resJSON, err := json.Marshal(responseData)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseCode{
            Code:    fiber.StatusInternalServerError,
            Message: "Failed to marshal book",
            Data:    nil,
        })
    }

    err = controller.RedisClient.Set(c.Context(), cacheKey, resJSON, 5 * time.Minute).Err()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseCode{
            Code:    fiber.StatusInternalServerError,
            Message: "Failed to cache book",
            Data:    nil,
        })
    }

    c.Set("X-Source", "database")
    log.Printf("FindById from database took %v", time.Since(startTime))
    return c.JSON(models.ResponseCode{
        Code:    fiber.StatusOK,
        Message: "Book found",
        Data:    responseData,
    })
}

func (controller *BooksController) Delete(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    err := controller.BooksService.Delete(c.Context(), id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(models.ResponseCode{
            Code:    fiber.StatusNotFound,
            Message: err.Error(),
            Data:    nil,
        })
    }

    // Clear the cache for the specific book
    cacheKey := os.Getenv("CACHE_KEY_BOOKS_PREFIX") + strconv.Itoa(id)
    controller.RedisClient.Del(c.Context(), cacheKey)

    // Clear the cache for books list after deleting a book
    controller.RedisClient.Del(c.Context(), os.Getenv("CACHE_KEY_BOOKS_ALL"))

    return c.JSON(models.ResponseCode{
        Code:    fiber.StatusNoContent,
        Message: "Book deleted successfully",
        Data:    nil,
    })
}
