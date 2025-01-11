package routes

import "github.com/gofiber/fiber/v2"

func AuthRoute(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Get("/get")

	auth.Post("/sign-up")
	auth.Post("/log-in")
	auth.Post("/reset-password")
}
