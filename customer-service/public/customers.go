package public

import (
	"errors"

	"github.com/Rx-11/EDIS-A2/customer-service/common"
	"github.com/Rx-11/EDIS-A2/customer-service/db"
	"github.com/Rx-11/EDIS-A2/customer-service/pkg"
	"github.com/Rx-11/EDIS-A2/customer-service/pkg/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func fetchUserById(c *fiber.Ctx) error {

	param := c.Locals("param").(fetchUserByIdParam)

	user, err := pkg.UserRepo.FetchUserByID(db.GetDB(), param.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(common.ErrNotFound.StatusCode).JSON(common.ErrNotFound)
		}
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	return c.JSON(user)
}

func fetchUserByUserId(c *fiber.Ctx) error {

	query := c.Locals("query").(fetchUserByUserIdQuery)

	user, err := pkg.UserRepo.FetchUserByUserID(db.GetDB(), query.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(common.ErrNotFound.StatusCode).JSON(common.ErrNotFound)
		}
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}

	return c.JSON(user)
}

func createUser(c *fiber.Ctx) error {
	body := c.Locals("body").(createUserRequest)

	existingUser, _ := pkg.UserRepo.FetchUserByUserID(db.GetDB(), body.UserID)
	if existingUser != nil {
		return c.Status(common.ErrUnprocessableEntity.StatusCode).JSON(common.NewError(
			common.ErrUnprocessableEntity.StatusCode,
			"This user ID already exists in the system.",
		))
	}

	user, err := pkg.UserRepo.CreateUser(db.GetDB(), models.User{
		UserID:   body.UserID,
		Name:     body.Name,
		Phone:    body.Phone,
		Address:  body.Address,
		Address2: body.Address2,
		City:     body.City,
		State:    body.State,
		Zipcode:  body.Zipcode,
	})
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).JSON(common.ErrInternalServerError)
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
