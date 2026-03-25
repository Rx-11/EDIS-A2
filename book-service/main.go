package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Rx-11/EDIS-A2/book-service/config"
	"github.com/Rx-11/EDIS-A2/book-service/db"
	"github.com/Rx-11/EDIS-A2/book-service/public"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "BookStore-Backend",
	})

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path}\nBody: ${body}\n",
	}))

	config.Init()
	log.Println("Loaded configs.")
	db.Init(config.GetConfig().DbConfig, db.MySQL, db.LogInfo)
	db.Migrate()

	public.MountRoutes(app)

	go func() {
		log.Println("Server started at http://localhost:3000")
		if err := app.Listen("0.0.0.0:3000"); err != nil {
			log.Printf("Server startup error: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	log.Println("Gracefully shutting down...")
	if err := app.Shutdown(); err != nil {
		log.Printf("Error shutting down server: %v", err)
	}
	log.Println("Server stopped")
}
