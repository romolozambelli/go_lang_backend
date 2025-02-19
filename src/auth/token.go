package auth

import (
	"backend/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Generate a token to authorize actions
func GenerateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID

	// secret for signing the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}
