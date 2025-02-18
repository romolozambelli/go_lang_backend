package security

import "golang.org/x/crypto/bcrypt"

// Converting a password in a hash code
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Compare hash and password to validate password
func CheckPassword(passwordHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}
