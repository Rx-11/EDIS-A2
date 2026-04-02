package public

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func MountRoutes(router *fiber.App) {
	router.Get("/", func(c *fiber.Ctx) error {
		log.Println("OK")
		return c.SendString("OK")
	})

	router.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	userGroup := router.Group("/customers")
	{
		userGroup.Post("/", jwtMiddleware(), requireClientType(), parseBody(createUserRequest{}), createUser)
		userGroup.Get("/", jwtMiddleware(), requireClientType(), parseQuery(fetchUserByUserIdQuery{}), fetchUserByUserId)
		userGroup.Get("/:id", jwtMiddleware(), requireClientType(), parseParam(fetchUserByIdParam{}), fetchUserById)
	}

	bookGroup := router.Group("/books")
	{
		bookGroup.Post("/", jwtMiddleware(), requireClientType(), parseBody(createBookRequest{}), createBook)
		bookGroup.Get("/isbn/:isbn", jwtMiddleware(), requireClientType(), parseParam(fetchBookByISBNParam{}), fetchBookByISBN)
		bookGroup.Get("/:isbn", jwtMiddleware(), requireClientType(), parseParam(fetchBookByISBNParam{}), fetchBookByISBN)
		bookGroup.Put("/:isbn", jwtMiddleware(), requireClientType(), parseParam(fetchBookByISBNParam{}), parseBody(updateBookRequest{}), updateBook)
	}

}
