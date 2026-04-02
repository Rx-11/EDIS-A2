package public

import (
	"log"

	"github.com/Rx-11/EDIS-A2/book-web-bff/common"
	"github.com/Rx-11/EDIS-A2/book-web-bff/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v3/client"
)

func fetchBookByISBN(c *fiber.Ctx) error {
	param := c.Locals("param").(fetchBookByISBNParam)

	targetURL := config.GetConfig().BookSvcURL + "/books/" + param.ISBN
	log.Printf("web-bff fetchBookByISBN -> GET %s", targetURL)

	resp, err := config.GetFiberClient().Get(targetURL)
	if err != nil {
		log.Printf("web-bff fetchBookByISBN downstream error: %v", err)
		return c.Status(common.ErrInternalServerError.StatusCode).
			JSON(common.ErrInternalServerError)
	}

	log.Printf("web-bff fetchBookByISBN <- status=%d body=%s", resp.StatusCode(), string(resp.Body()))
	return c.Status(resp.StatusCode()).Send(resp.Body())
}

func createBook(c *fiber.Ctx) error {
	body := c.Locals("body").(createBookRequest)

	targetURL := config.GetConfig().BookSvcURL + "/books"
	log.Printf("web-bff createBook -> POST %s body=%+v", targetURL, body)

	resp, err := config.GetFiberClient().Post(targetURL, client.Config{Body: body})
	if err != nil {
		log.Printf("web-bff createBook downstream error: %v", err)
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	log.Printf("web-bff createBook <- status=%d body=%s", resp.StatusCode(), string(resp.Body()))
	return c.Status(resp.StatusCode()).Send(resp.Body())
}

func updateBook(c *fiber.Ctx) error {
	param := c.Locals("param").(fetchBookByISBNParam)
	body := c.Locals("body").(updateBookRequest)

	targetURL := config.GetConfig().BookSvcURL + "/books/" + param.ISBN
	log.Printf("web-bff updateBook -> PUT %s body=%+v", targetURL, body)

	resp, err := config.GetFiberClient().Put(
		targetURL,
		client.Config{Body: body},
	)
	if err != nil {
		log.Printf("web-bff updateBook downstream error: %v", err)
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	log.Printf("web-bff updateBook <- status=%d body=%s", resp.StatusCode(), string(resp.Body()))
	return c.Status(resp.StatusCode()).Send(resp.Body())
}
