package main

import (
	"basic-jwt-auth/config"
	"basic-jwt-auth/internal/handlers"
	"basic-jwt-auth/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Middleware logging
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Create a new JWT middleware
	// Note: This is just an example, please use a secure secret key
	jwt := middlewares.NewAuthMiddleware(config.Secret)

	// Create a Login route
	app.Post("/login", handlers.Login)

	// Create a protected route
	app.Get("/protected", jwt, handlers.Protected)

	// Listen on port 3000
	err := app.Listen(":3003")
	if err != nil {
		return
	}
}
