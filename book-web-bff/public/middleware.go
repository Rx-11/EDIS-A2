package public

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Rx-11/EDIS-A2/book-web-bff/common"
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

func parseBody[T any](_ T) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body T
		if err := c.BodyParser(&body); err != nil {
			log.Printf("parseBody error path=%s raw=%s err=%v", c.Path(), string(c.Body()), err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Invalid params - body",
			})
		}

		if err := validate.Struct(body); err != nil {
			log.Printf("parseBody validation failed path=%s body=%+v err=%v", c.Path(), body, err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Validation failed - body",
			})
		}

		log.Printf("parseBody ok path=%s body=%+v", c.Path(), body)
		c.Locals("body", body)
		return c.Next()
	}
}

func parseQuery[T any](_ T) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query T
		if err := c.QueryParser(&query); err != nil {
			log.Printf("parseQuery error path=%s rawQuery=%s err=%v", c.Path(), c.Context().QueryArgs().String(), err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Invalid params - query",
			})
		}

		if err := validate.Struct(query); err != nil {
			log.Printf("parseQuery validation failed path=%s rawQuery=%s parsed=%+v err=%v", c.Path(), c.Context().QueryArgs().String(), query, err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Validation failed - query",
			})
		}

		log.Printf("parseQuery ok path=%s parsed=%+v", c.Path(), query)
		c.Locals("query", query)
		return c.Next()
	}
}

func parseParam[T any](_ T) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var param T
		if err := c.ParamsParser(&param); err != nil {
			log.Printf("parseParam error path=%s params=%v err=%v", c.Path(), c.AllParams(), err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Invalid params - param",
			})
		}

		if err := validate.Struct(param); err != nil {
			log.Printf("parseParam validation failed path=%s params=%v parsed=%+v err=%v", c.Path(), c.AllParams(), param, err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Validation failed - param",
			})
		}

		log.Printf("parseParam ok path=%s parsed=%+v", c.Path(), param)
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

func requireClientType() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client := c.Get("X-Client-Type")
		if client == "" {
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Missing X-Client-Type header",
			})
		}

		allowed := map[string]bool{
			"web":     true,
			"ios":     true,
			"android": true,
		}
		if !allowed[strings.ToLower(client)] {
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Invalid X-Client-Type header",
			})
		}

		c.Locals("target", client)
		return c.Next()
	}
}

func jwtMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or invalid token",
			})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		parts := strings.Split(tokenStr, ".")
		if len(parts) != 3 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token format",
			})
		}

		payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token payload",
			})
		}

		var claims map[string]interface{}
		if err := json.Unmarshal(payloadBytes, &claims); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token payload",
			})
		}

		validSubs := map[string]bool{
			"starlord": true,
			"gamora":   true,
			"drax":     true,
			"rocket":   true,
			"groot":    true,
		}

		sub, ok := claims["sub"].(string)
		if !ok || !validSubs[sub] {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token subject",
			})
		}

		iss, ok := claims["iss"].(string)
		if !ok || iss != "cmu.edu" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token issuer",
			})
		}

		expVal, ok := claims["exp"]
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing token expiration",
			})
		}

		expFloat, ok := expVal.(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token expiration",
			})
		}

		if int64(expFloat) <= time.Now().Unix() {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token expired",
			})
		}

		return c.Next()
	}
}
