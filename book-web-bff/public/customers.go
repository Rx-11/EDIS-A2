package public

import (
	"log"
	"net/url"
	"strconv"

	"github.com/Rx-11/EDIS-A2/book-web-bff/common"
	"github.com/Rx-11/EDIS-A2/book-web-bff/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v3/client"
)

func fetchUserById(c *fiber.Ctx) error {
	param := c.Locals("param").(fetchUserByIdParam)

	targetURL := config.GetConfig().CustomerSvcURL + "/customers/" + strconv.Itoa(int(param.ID))
	log.Printf("web-bff fetchUserById -> GET %s", targetURL)

	resp, err := config.GetFiberClient().Get(targetURL)
	if err != nil {
		log.Printf("web-bff fetchUserById downstream error: %v", err)
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	log.Printf("web-bff fetchUserById <- status=%d body=%s", resp.StatusCode(), string(resp.Body()))
	return c.Status(resp.StatusCode()).Send(resp.Body())
}

func fetchUserByUserId(c *fiber.Ctx) error {
	query := c.Locals("query").(fetchUserByUserIdQuery)

	targetURL := config.GetConfig().CustomerSvcURL + "/customers?userId=" + url.QueryEscape(query.UserID)
	log.Printf("web-bff fetchUserByUserId -> GET %s", targetURL)

	resp, err := config.GetFiberClient().Get(targetURL)
	if err != nil {
		log.Printf("web-bff fetchUserByUserId downstream error: %v", err)
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	log.Printf("web-bff fetchUserByUserId <- status=%d body=%s", resp.StatusCode(), string(resp.Body()))
	return c.Status(resp.StatusCode()).Send(resp.Body())
}

func createUser(c *fiber.Ctx) error {
	body := c.Locals("body").(createUserRequest)

	targetURL := config.GetConfig().CustomerSvcURL + "/customers"
	log.Printf("web-bff createUser -> POST %s body=%+v", targetURL, body)

	resp, err := config.GetFiberClient().Post(targetURL, client.Config{Body: body})
	if err != nil {
		log.Printf("web-bff createUser downstream error: %v", err)
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	log.Printf("web-bff createUser <- status=%d body=%s", resp.StatusCode(), string(resp.Body()))
	return c.Status(resp.StatusCode()).Send(resp.Body())
}
