package middlewares

import (
	"golang_basic_gin_sept_2023/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Request need access token",
				"status":  http.StatusUnauthorized,
			})

			c.Abort()
			return
		}

		//:: VALIDATE TOKEN
		email, role, err := config.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
				"status":  http.StatusUnauthorized,
			})

			c.Abort()
			return
		}

		//:: SET EMAIL
		c.Set("x-email", email)
		c.Set("x-role", role)

		//:: NEXT
		c.Next()
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Request need access token",
				"status":  http.StatusUnauthorized,
			})

			c.Abort()
			return
		}

		//:: VALIDATE TOKEN
		email, role, err := config.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
				"status":  http.StatusUnauthorized,
			})

			c.Abort()
			return
		}

		if role != 1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Your role is't access this end point",
				"status":  http.StatusUnauthorized,
			})

			c.Abort()
			return
		}

		//:: SET EMAIL
		c.Set("x-email", email)
		c.Set("x-role", role)

		//:: NEXT
		c.Next()
	}
}
