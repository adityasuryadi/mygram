package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ClaimToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

func DecodeToken() (jwt.MapClaims, bool) {
	var token *jwt.Token
	claims, err := token.Claims.(jwt.MapClaims)
	if err {
		return nil, true
	}
	return claims, false
}

