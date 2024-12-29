package main

import (
	"fmt"
	"github.com/ArtiomStartev/jwt-auth-api/database"
	"github.com/ArtiomStartev/jwt-auth-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	if err := database.DBConn(); err != nil {
		fmt.Println("Error connecting to database: ", err)
		return
	}

	routes.Setup(app)

	if err := app.Listen(":8000"); err != nil {
		fmt.Println("Error starting server on port 8000: ", err)
		return
	}
}
