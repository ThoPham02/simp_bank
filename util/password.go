package util

import "golang.org/x/crypto/bcrypt"

// HashPassword will return the password associated
func HashPassword(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(byte), nil
}

// ComparePassword will return true if the password is the same
func ComparePassword(password string, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
