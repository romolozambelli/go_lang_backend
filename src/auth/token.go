package auth

import (
	"backend/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

// Function to check if the given token is valid or not
func CheckToken(r *http.Request) error {
	tokenString := getToken(r)
	token, erro := jwt.Parse(tokenString, checktSigningMethod)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("invalid token")

}

// Get the userID included on the token
func GetUserIDFromToken(r *http.Request) (uint64, error) {

	tokenString := getToken(r)
	token, erro := jwt.Parse(tokenString, checktSigningMethod)
	if erro != nil {
		return 0, erro
	}

	if perm, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", perm["userID"]), 10, 64)

		if erro != nil {
			return 0, erro
		}
		return userID, nil
	}

	return 0, errors.New("invalid token")

}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""

}

func checktSigningMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("wrong signature method ! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
