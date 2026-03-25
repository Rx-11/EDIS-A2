package public

import (
	"errors"
	"fmt"

	"github.com/Rx-11/EDIS-A2/book-service/ai"
	"github.com/Rx-11/EDIS-A2/book-service/common"
	"github.com/Rx-11/EDIS-A2/book-service/config"
	"github.com/Rx-11/EDIS-A2/book-service/db"
	"github.com/Rx-11/EDIS-A2/book-service/pkg"
	"github.com/Rx-11/EDIS-A2/book-service/pkg/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func fetchBookByISBN(c *fiber.Ctx) error {

	param := c.Locals("param").(fetchBookByISBNParam)

	book, err := pkg.BookRepo.FetchBookByISBN(db.GetDB(), param.ISBN)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(common.ErrNotFound.StatusCode).JSON(common.ErrNotFound)
		}
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	return c.JSON(book)
}

func createBook(c *fiber.Ctx) error {
	body := c.Locals("body").(createBookRequest)

	existingBook, _ := pkg.BookRepo.FetchBookByISBN(db.GetDB(), body.ISBN)
	if existingBook != nil {
		return c.Status(common.ErrUnprocessableEntity.StatusCode).JSON(common.NewError(
			common.ErrUnprocessableEntity.StatusCode,
			"This ISBN already exists in the system.",
		))
	}

	book, err := pkg.BookRepo.CreateBook(db.GetDB(), models.Book{
		ISBN:        body.ISBN,
		Title:       body.Title,
		Author:      body.Author,
		Price:       body.Price,
		Description: body.Description,
		Genre:       body.Genre,
		Quantity:    *body.Quantity,
	})
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	SummaryResp, err := config.Gemini.Chat(ai.ChatRequest{Messages: []ai.Message{{Role: "model", Content: "Give a 500 word summary of the following book"}, {Role: "user", Content: fmt.Sprintf("Book Title: %s\nBook Description: %s\nBook Author: %s\nBook ISBN: %s", book.Title, book.Description, book.Author, book.ISBN)}}})
	if err == nil && SummaryResp.Response != "" {
		book.Summary = &SummaryResp.Response
		pkg.BookRepo.UpdateBook(db.GetDB(), *book)
	}
	responseBook := *book
	responseBook.Summary = nil

	return c.Status(fiber.StatusCreated).JSON(responseBook)
}

func updateBook(c *fiber.Ctx) error {
	param := c.Locals("param").(fetchBookByISBNParam)
	body := c.Locals("body").(updateBookRequest)

	if param.ISBN != body.ISBN {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ISBN in URL does not match ISBN in body"})
	}

	existingBook, err := pkg.BookRepo.FetchBookByISBN(db.GetDB(), param.ISBN)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(common.ErrNotFound.StatusCode).JSON(common.ErrNotFound)
		}
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	existingBook.Title = body.Title
	existingBook.Author = body.Author
	existingBook.Price = body.Price
	existingBook.Description = body.Description
	existingBook.Genre = body.Genre
	existingBook.Quantity = *body.Quantity

	book, err := pkg.BookRepo.UpdateBook(db.GetDB(), *existingBook)
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}
	book.Summary = nil
	return c.Status(fiber.StatusOK).JSON(book)
}
