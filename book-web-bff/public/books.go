package public

import (
	"github.com/Rx-11/EDIS-A2/book-web-bff/common"
	"github.com/Rx-11/EDIS-A2/book-web-bff/config"
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

	resp, err := config.GetFiberClient().Put(
		config.GetConfig().BookSvcURL+"/books/"+param.ISBN,
		client.Config{Body: body},
	)
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	return c.Status(resp.StatusCode()).Send(resp.Body())
}
