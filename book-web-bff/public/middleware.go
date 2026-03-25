package public

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Rx-11/EDIS-A2/book-web-bff/common"
	"github.com/Rx-11/EDIS-A2/book-web-bff/config"
	"github.com/golang-jwt/jwt/v4"
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
		err := c.BodyParser(&body)
		if err != nil {
			log.Println(err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Invalid params - body",
			})
		}

		err = validate.Struct(body)
		if err != nil {
			log.Println("Validation error:", err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Validation failed - body",
			})
		}
		c.Locals("body", body)
		return c.Next()
	}
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

func parseParam[T any](_ T) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var param T
		err := c.ParamsParser(&param)
		if err != nil {
			log.Println(err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Invalid params - param",
			})
		}

		err = validate.Struct(param)
		if err != nil {
			log.Println("Validation error:", err)
			return c.Status(common.ErrInvalidParams.StatusCode).JSON(fiber.Map{
				"error": "Validation failed - param",
			})
		}
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
			"Web":     true,
			"iOS":     true,
			"Android": true,
		}
		if !allowed[client] {
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
		tokenStr := ""

		authHeader := c.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenStr = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			tokenStr = c.Cookies("token")
		}

		if tokenStr == "" {
			return c.Status(common.ErrUnauthorized.StatusCode).JSON(fiber.Map{
				"error": "Missing or invalid token",
			})
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(config.GetConfig().ACSecrets.UserSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(common.ErrUnauthorized.StatusCode).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(common.ErrUnauthorized.StatusCode).JSON(fiber.Map{
				"error": "Invalid token payload",
			})
		}

		iss, ok := claims["iss"].(string)
		if !ok || iss != "cmu.edu" {
			return c.Status(common.ErrUnauthorized.StatusCode).JSON(fiber.Map{
				"error": "Invalid token issuer",
			})
		}

		sub, ok := claims["sub"].(string)
		if !ok {
			return c.Status(common.ErrUnauthorized.StatusCode).JSON(fiber.Map{
				"error": "Invalid token subject",
			})
		}
		allowed := map[string]bool{
			"starlord": true,
			"gamora":   true,
			"drax":     true,
			"rocket":   true,
			"groot":    true,
		}
		if !allowed[sub] {
			return c.Status(common.ErrUnauthorized.StatusCode).JSON(fiber.Map{
				"error": "Unauthorized subject",
			})
		}

		expVal, ok := claims["exp"]
		if !ok {
			return c.Status(common.ErrUnauthorized.StatusCode).JSON(fiber.Map{
				"error": "Missing exp claim",
			})
		}
		var expInt int64
		switch v := expVal.(type) {
		case float64:
			expInt = int64(v)
		case int64:
			expInt = v
		case json.Number:
			n, err := v.Int64()
			if err != nil {
				return c.Status(common.ErrUnauthorized.StatusCode).JSON(fiber.Map{
					"error": "Invalid exp claim",
				})
			}
			expInt = n
		default:
			return c.Status(common.ErrUnauthorized.StatusCode).JSON(fiber.Map{
				"error": "Invalid exp claim type",
			})
		}
		if time.Unix(expInt, 0).Before(time.Now()) {
			return c.Status(common.ErrUnauthorized.StatusCode).JSON(fiber.Map{
				"error": "Token expired",
			})
		}

		c.Locals("user_id", sub)
		return c.Next()
	}
}
