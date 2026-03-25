package public

import (
	"encoding/json"

	"github.com/Rx-11/EDIS-A2/book-mobile-bff/common"
	"github.com/Rx-11/EDIS-A2/book-mobile-bff/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v3/client"
)

func fetchBookByISBN(c *fiber.Ctx) error {
	param := c.Locals("param").(fetchBookByISBNParam)

	url := config.GetConfig().BookSvcURL + "/books/" + param.ISBN

	resp, err := config.GetFiberClient().Get(url)
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).
			JSON(common.ErrInternalServerError)
	}

	if resp.StatusCode() != fiber.StatusOK {
		return c.Status(resp.StatusCode()).Send(resp.Body())
	}

	var bookResponse bookResponse
	err = json.Unmarshal(resp.Body(), &bookResponse)
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).
			JSON(common.ErrInternalServerError)
	}

	if bookResponse.Genre == "non fiction" {
		book := getbookResponse{
			ISBN:        bookResponse.ISBN,
			Title:       bookResponse.Title,
			Author:      bookResponse.Author,
			Genre:       3,
			Price:       bookResponse.Price,
			Description: bookResponse.Description,
			Quantity:    bookResponse.Quantity,
		}
		return c.JSON(book)
	}

	c.Status(resp.StatusCode())

	return c.Send(resp.Body())
}

func createBook(c *fiber.Ctx) error {
	body := c.Locals("body").(createBookRequest)

	resp, err := config.GetFiberClient().Post(config.GetConfig().BookSvcURL+"/books/", client.Config{Body: body})
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	return c.Status(resp.StatusCode()).Send(resp.Body())
}

func updateBook(c *fiber.Ctx) error {
	param := c.Locals("param").(fetchBookByISBNParam)
	body := c.Locals("body").(updateBookRequest)

	resp, err := config.GetFiberClient().Put(config.GetConfig().BookSvcURL+"/books/"+param.ISBN, client.Config{Body: body})
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(resp.Body())
}
