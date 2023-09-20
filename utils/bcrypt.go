package utils

import "golang.org/x/crypto/bcrypt"

func HashString(password string) []byte {
	hashedString, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return hashedString
}

func CompareHashAndPassword(hashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return false
	}
	return true
}
