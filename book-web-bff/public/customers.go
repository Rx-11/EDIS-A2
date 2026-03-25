package public

import (
	"strconv"

	"github.com/Rx-11/EDIS-A2/book-web-bff/common"
	"github.com/Rx-11/EDIS-A2/book-web-bff/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v3/client"
)

func fetchUserById(c *fiber.Ctx) error {

	param := c.Locals("param").(fetchUserByIdParam)

	resp, err := config.GetFiberClient().Get(config.GetConfig().CustomerSvcURL + "/customers/" + strconv.Itoa(int(param.ID)))
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	return c.Status(resp.StatusCode()).Send(resp.Body())
}

func fetchUserByUserId(c *fiber.Ctx) error {

	query := c.Locals("query").(fetchUserByUserIdQuery)

	resp, err := config.GetFiberClient().Get(config.GetConfig().CustomerSvcURL + "/customers?user_id=" + query.UserID)
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	return c.Status(resp.StatusCode()).Send(resp.Body())

}

func createUser(c *fiber.Ctx) error {
	body := c.Locals("body").(createUserRequest)

	resp, err := config.GetFiberClient().Post(config.GetConfig().CustomerSvcURL+"/customers/", client.Config{Body: body})
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	return c.Status(resp.StatusCode()).Send(resp.Body())
}
