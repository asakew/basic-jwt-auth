package repository

import (
	"errors"

	"basic-jwt-auth/internal/models"
)

// Simulate a database call
func FindByCredentials(login, password string) (*models.User, error) {
	// Here you would query your database for the user with the given email
	if login == "test12345" && password == "test12345" {
		return &models.User{
			ID:             1,
			Login:          "test12345",
			Password:       "test12345",
			FavoritePhrase: "Hello, World!",
		}, nil
	}

	return nil, errors.New("user not found")
}
