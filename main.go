package main

import (
	"github.com/dogayaglicioglu/go-oauth2/config"
	"github.com/dogayaglicioglu/go-oauth2/controllers"
	"github.com/dogayaglicioglu/go-oauth2/endpoints"
	"github.com/dogayaglicioglu/go-oauth2/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.GoogleConfig()

	app.Get("/google_login", controllers.GoogleLogin)
	app.Get("/google_callback", controllers.GoogleCallback)

	protectedRoutes := app.Group("/api/protected", middleware.Middleware)
	protectedRoutes.Get("/example", endpoints.ProtectedEndpoint)
	app.Listen(":8080")

}
