package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/killerrekt/Go-Fiber-Auth/internal/controller"
	"github.com/killerrekt/Go-Fiber-Auth/internal/middleware"
)

func AuthRoute(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Get("/get", middleware.AuthenticateAndAuthorize(), controller.Me)

	auth.Post("/sign-up", controller.SignUp)
	auth.Post("/log-in", controller.LogIn)
	auth.Post("/reset-password")
}
