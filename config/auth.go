package config

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("rahasia")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     uint32 `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(email, username string, role uint32) (tokenString string, err error) {
	expTime := time.Now().Add(1 * time.Hour)
	// Replace "your-secret-key" with a strong, secret key

	claims := &JWTClaim{
		Email:    email,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)

	return
}
