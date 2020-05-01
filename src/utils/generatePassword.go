package utils

import (
	"math/rand"
)

var DefaultCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func GeneratePassword() string {
	var password string

	for i := 0; i < 8; i++ {
		password += string(DefaultCharacters[rand.Intn(len(DefaultCharacters))])
	}

	return password
}
