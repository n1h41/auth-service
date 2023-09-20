package utils

import (
	"n1h41/auth-service/models"
	"testing"
)

func TestCreateJwtToken(t *testing.T) {
	user := models.UserModel{
		ID:        1,
		Email:     "testuser@gmail.com",
		Name:      "Test User",
		Password:  "testpassword",
		CreatedAt: "2021-01-01",
		UpdatedAt: "2021-01-01",
	}
	token, err := CreateJwtToken(user)
	if err != nil {
		t.Fatalf("Error creating jwt token: %s", err)
	}
	println(token)
}
