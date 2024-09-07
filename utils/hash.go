package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hPass), err
}
