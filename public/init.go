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
		userGroup.Post("/", parseBody(createUserRequest{}), createUser)
		userGroup.Get("/", parseQuery(fetchUserByUserIdQuery{}), fetchUserByUserId)
		userGroup.Get("/:id", parseParam(fetchUserByIdParam{}), fetchUserById)
	}

	bookGroup := router.Group("/books")
	{
		bookGroup.Post("/", parseBody(createBookRequest{}), createBook)
		bookGroup.Get("/isbn/:isbn", parseParam(fetchBookByISBNParam{}), fetchBookByISBN)
		bookGroup.Get("/:isbn", parseParam(fetchBookByISBNParam{}), fetchBookByISBN)
		bookGroup.Put("/:isbn", parseParam(fetchBookByISBNParam{}), parseBody(updateBookRequest{}), updateBook)
	}

}
