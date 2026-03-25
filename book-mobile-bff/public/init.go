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
		userGroup.Post("/", requireClientType(), jwtMiddleware(), parseBody(createUserRequest{}), createUser)
		userGroup.Get("/", requireClientType(), jwtMiddleware(), parseQuery(fetchUserByUserIdQuery{}), fetchUserByUserId)
		userGroup.Get("/:id", requireClientType(), jwtMiddleware(), parseParam(fetchUserByIdParam{}), fetchUserById)
	}

	bookGroup := router.Group("/books")
	{
		bookGroup.Post("/", requireClientType(), jwtMiddleware(), parseBody(createBookRequest{}), createBook)
		bookGroup.Get("/isbn/:isbn", requireClientType(), jwtMiddleware(), parseParam(fetchBookByISBNParam{}), fetchBookByISBN)
		bookGroup.Get("/:isbn", requireClientType(), jwtMiddleware(), parseParam(fetchBookByISBNParam{}), fetchBookByISBN)
		bookGroup.Put("/:isbn", requireClientType(), jwtMiddleware(), parseParam(fetchBookByISBNParam{}), parseBody(updateBookRequest{}), updateBook)
	}

}
