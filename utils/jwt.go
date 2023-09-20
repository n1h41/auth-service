package utils

import (
	// "fmt"
	// "n1h41/auth-service/config"
	"fmt"
	"n1h41/auth-service/config"
	"n1h41/auth-service/models"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("secret")

func CreateJwtToken(user models.UserModel) (string, error) {
	config, _ := config.LoadConfig("../")
  fmt.Println("JWT SECRET:  " + config.JwtSecret)
	jwtSecret := []byte(config.JwtSecret)
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
