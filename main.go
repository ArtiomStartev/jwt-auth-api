package main

import (
	"fmt"
	"github.com/ArtiomStartev/jwt-auth-api/database"
	"github.com/ArtiomStartev/jwt-auth-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DBConn()

	app := fiber.New()

	routes.Setup(app)

	if err := app.Listen(":8000"); err != nil {
		fmt.Println("Error listening http requests on port :8000", err)
		return
	}
}
