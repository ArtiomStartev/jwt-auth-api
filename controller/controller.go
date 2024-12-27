package controller

import (
	"fmt"
	"github.com/ArtiomStartev/jwt-auth-api/database"
	"github.com/ArtiomStartev/jwt-auth-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const SecretKey = "secret" // Secret key for JWT Token

func User(ctx *fiber.Ctx) error {
	var user models.User
	cookie := ctx.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil // use the same secret key that was used to sign the token
	})

	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"data":  nil,
			"error": "Unauthorized",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	if err = database.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		fmt.Println("Error fetching user: ", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"data":  nil,
			"error": "User not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  user,
		"error": nil,
	})
}

func Register(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		fmt.Println("Error parsing request body: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":  nil,
			"error": "Error parsing request body",
		})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error generating password hash: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":  nil,
			"error": "Error registering user",
		})
	}

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	if err = database.DB.Create(&user).Error; err != nil {
		fmt.Println("Error creating user: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":  nil,
			"error": "Error registering user",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":  user,
		"error": nil,
	})
}

func Login(ctx *fiber.Ctx) error {
	var user models.User
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		fmt.Println("Error parsing request body: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":  nil,
			"error": "Error parsing request body",
		})
	}

	// Check if a user exists in the database by email
	if err := database.DB.Where("email = ?", data["email"]).First(&user).Error; err != nil {
		fmt.Println("Error fetching user: ", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"data":  nil,
			"error": "User not found",
		})
	}

	if user.ID == 0 {
		ctx.Status(fiber.StatusNotFound)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":  nil,
			"error": "User not found",
		})
	}

	// Compare password hash with the password provided by the user
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":  nil,
			"error": "Invalid password",
		})
	}

	// Create JWT Token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),            // issuer contains the ID of the user
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // JWT Token expires in 24 hours
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":  nil,
			"error": "Error logging in",
		})
	}

	// Set JWT Token in a cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	ctx.Cookie(&cookie)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  user,
		"error": nil,
	})
}

func Logout(ctx *fiber.Ctx) error {
	// Clear the JWT Token cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	ctx.Cookie(&cookie)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  nil,
		"error": nil,
	})
}
