package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashFromPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashFromPassword), []byte(password))
}
