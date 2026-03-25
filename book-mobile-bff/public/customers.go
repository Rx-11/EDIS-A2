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

	resp, err := config.GetFiberClient().Get(config.GetConfig().CustomerSvcURL + "/customers/" + strconv.Itoa(int(param.ID)))
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	var user userResponse
	err = json.Unmarshal(resp.Body(), &user)
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}
	userResp := getUserResponse{
		ID:     user.ID,
		UserID: user.UserID,
		Name:   user.Name,
		Phone:  user.Phone,
	}
	if resp.StatusCode() != fiber.StatusOK {
		return c.Status(resp.StatusCode()).Send(resp.Body())
	}

	return c.JSON(userResp)
}

func fetchUserByUserId(c *fiber.Ctx) error {

	query := c.Locals("query").(fetchUserByUserIdQuery)

	resp, err := config.GetFiberClient().Get(config.GetConfig().CustomerSvcURL + "/customers?user_id=" + query.UserID)
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	var user userResponse
	err = json.Unmarshal(resp.Body(), &user)
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}
	userResp := getUserResponse{
		ID:     user.ID,
		UserID: user.UserID,
		Name:   user.Name,
		Phone:  user.Phone,
	}
	if resp.StatusCode() != fiber.StatusOK {
		return c.Status(resp.StatusCode()).Send(resp.Body())
	}

	return c.JSON(userResp)

}

func createUser(c *fiber.Ctx) error {
	body := c.Locals("body").(createUserRequest)

	resp, err := config.GetFiberClient().Post(config.GetConfig().CustomerSvcURL+"/customers/", client.Config{Body: body})
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	return c.Status(resp.StatusCode()).Send(resp.Body())
}
