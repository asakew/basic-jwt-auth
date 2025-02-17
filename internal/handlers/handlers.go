package handlers

import (
	"time"

	"basic-jwt-auth/config"
	"basic-jwt-auth/internal/models"
	"basic-jwt-auth/internal/repository"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	// Extract the credentials from the request body
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Find the user by credentials
	user, err := repository.FindByCredentials(loginRequest.Login, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	day := time.Hour * 24

	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"login": user.Login,
		"fav":   user.FavoritePhrase,
		"exp":   time.Now().Add(day * 1).Unix(),
	}

	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Return the token
	return c.JSON(models.LoginResponse{
		Token: t,
	})
}

// create html templates
func Protected(c *fiber.Ctx) error {
	// Get the user from the context and return it
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	login := claims["Login"].(string)
	favPhrase := claims["fav"].(string)
	return c.SendString("Welcome 👋" + login + " " + favPhrase)
}
