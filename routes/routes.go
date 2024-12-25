package routes

import (
	"github.com/ArtiomStartev/jwt-auth-api/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/user")

	api.Get("/get-user", controller.User)

	api.Post("/register", controller.Register)

	api.Post("/login", controller.Login)

	api.Post("/logout", controller.Logout)
}
