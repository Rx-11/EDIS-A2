package public

import (
	"log"
	"strconv"
	"strings"

	"github.com/Rx-11/EDIS-A2/book-service/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func init() {
	validate.RegisterValidation("decimals2", func(fl validator.FieldLevel) bool {
		val := fl.Field().Float()
		str := strconv.FormatFloat(val, 'f', -1, 64)
		if idx := strings.Index(str, "."); idx != -1 {
			return len(str)-idx-1 <= 2
		}
		return true
	})
}

func parseQuery[T any](_ T) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query T
		err := c.QueryParser(&query)
		if err != nil {
			log.Println(err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Invalid params - query",
			})
		}

		err = validate.Struct(query)
		if err != nil {
			log.Println("Validation error:", err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Validation failed - query",
			})
		}
		c.Locals("query", query)
		return c.Next()
	}
}

func parseBody[T any](_ T) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body T
		if err := c.BodyParser(&body); err != nil {
			log.Printf("book-service parseBody error path=%s raw=%s err=%v", c.Path(), string(c.Body()), err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Invalid params - body",
			})
		}

		if err := validate.Struct(body); err != nil {
			log.Printf("book-service parseBody validation failed path=%s body=%+v err=%v", c.Path(), body, err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Validation failed - body",
			})
		}

		log.Printf("book-service parseBody ok path=%s body=%+v", c.Path(), body)
		c.Locals("body", body)
		return c.Next()
	}
}

func parseParam[T any](_ T) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var param T
		if err := c.ParamsParser(&param); err != nil {
			log.Printf("book-service parseParam error path=%s params=%v err=%v", c.Path(), c.AllParams(), err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Invalid params - param",
			})
		}

		if err := validate.Struct(param); err != nil {
			log.Printf("book-service parseParam validation failed path=%s params=%v parsed=%+v err=%v", c.Path(), c.AllParams(), param, err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Validation failed - param",
			})
		}

		log.Printf("book-service parseParam ok path=%s parsed=%+v", c.Path(), param)
		c.Locals("param", param)
		return c.Next()
	}
}

func setPagination() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query paginationQueryStruct
		err := c.QueryParser(&query)
		if err != nil {
			log.Println(err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Invalid params - pagination",
			})
		}

		err = validate.Struct(query)
		if err != nil {
			log.Println("Validation error:", err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Validation failed - pagination",
			})
		}

		pagination := paginationStruct{
			Limit:  query.PerPage,
			Offset: (query.Page - 1) * query.PerPage,
		}
		if query.PerPage == 0 && query.Page == 0 {
			pagination = paginationStruct{
				Limit:  10,
				Offset: 0,
			}
		}

		c.Locals("pagination", pagination)
		return c.Next()
	}
}
