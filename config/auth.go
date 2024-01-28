package config

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("rahasia")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     uint32 `json:"role"`
	// jwt.RegisteredClaims
	jwt.StandardClaims
}

func GenerateJWT(email, username string, role uint32) (tokenString string, err error) {
	// expTime := time.Now().Add(1 * time.Hour)
	expTime := time.Now().Add(1 * time.Minute)

	claims := &JWTClaim{
		Email:    email,
		Username: username,
		Role:     role,
		// RegisteredClaims: jwt.RegisteredClaims{
		// 	ExpiresAt: jwt.NewNumericDate(expTime),
		// },
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)

	return
}

//:: AuthMiddleware is a simple middleware to check if the request has a valid token.
//:: VALIDATE JWT MY VERSION
func ValidateJWT() gin.HandlerFunc {
	//you should set it on .env file then call it from there
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

		// Parse the token and verify its signature
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check that the signing method is what we expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Printf("unexpected signing method: %v", token.Header["alg"])
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Return the secret key
			return jwtKey, nil
		})

		// Check for errors
		if err != nil {
			//http.StatusUnauthorized == 401
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   http.StatusUnauthorized,
			})
			c.Abort()
		}
		c.Next()
	}
}

//:: VALIDATE JWT COACH HERU VERSION
func ValidateToken(signedToken string) (email string, role uint32, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)

	//:: IF FAILED CLAIMS
	if !ok {
		err = errors.New("could Not parse claims from token")
		return
	}

	//:: IF EXPIRED TOKEN
	if claims.ExpiresAt < time.Now().Unix() {
		err = errors.New("token expired")
		return
	}

	role = claims.Role
	email = claims.Email
	return
}
