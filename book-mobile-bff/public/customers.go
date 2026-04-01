package public

import (
	"encoding/json"
	"strconv"

	"github.com/Rx-11/EDIS-A2/book-mobile-bff/common"
	"github.com/Rx-11/EDIS-A2/book-mobile-bff/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v3/client"
)

func fetchUserById(c *fiber.Ctx) error {
	param := c.Locals("param").(fetchUserByIdParam)

	resp, err := config.GetFiberClient().Get(
		config.GetConfig().CustomerSvcURL + "/customers/" + strconv.Itoa(int(param.ID)),
	)
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	if resp.StatusCode() != fiber.StatusOK {
		return c.Status(resp.StatusCode()).Send(resp.Body())
	}

	var user userResponse
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	userResp := getUserResponse{
		ID:     user.ID,
		UserID: user.UserID,
		Name:   user.Name,
		Phone:  user.Phone,
	}

	return c.Status(fiber.StatusOK).JSON(userResp)
}

func fetchUserByUserId(c *fiber.Ctx) error {
	query := c.Locals("query").(fetchUserByUserIdQuery)

	resp, err := config.GetFiberClient().Get(
		config.GetConfig().CustomerSvcURL + "/customers?userId=" + query.UserID,
	)
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	if resp.StatusCode() != fiber.StatusOK {
		return c.Status(resp.StatusCode()).Send(resp.Body())
	}

	var user userResponse
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	userResp := getUserResponse{
		ID:     user.ID,
		UserID: user.UserID,
		Name:   user.Name,
		Phone:  user.Phone,
	}

	return c.Status(fiber.StatusOK).JSON(userResp)
}

func createUser(c *fiber.Ctx) error {
	body := c.Locals("body").(createUserRequest)

	resp, err := config.GetFiberClient().Post(config.GetConfig().CustomerSvcURL+"/customers/", client.Config{Body: body})
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	return c.Status(resp.StatusCode()).Send(resp.Body())
}
