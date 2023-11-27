package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewFiberApp(handler Handler) (*fiber.App, func(), error) {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	apiV1 := app.Group("/v1")
	apiV1.Post("/game/create", handler.CreateGame())

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	return app, func() {
		app.Shutdown()
	}, nil
}
