package routes

import (
	"golang_basic_gin_sept_2023/config"
	"golang_basic_gin_sept_2023/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})

		c.Abort()
		return
	}

	//:: HAS USER PASSWORD
	err := user.HashPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})

		c.Abort()
		return
	}

	//:: INSERT USER
	insertUser := config.DB.Create(&user)

	if insertUser.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error insert user",
			"error":   insertUser.Error.Error(),
		})

		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user_id":  user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
}

func GenerateToken(c *gin.Context) {
	request := models.TokenRequest{}
	user := models.User{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})

		c.Abort()
		return
	}

	//:: CHECK EMAIL
	checkEmail := config.DB.Where("email = ?", request.Email).First(&user)

	if checkEmail.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Email not found",
			"error":   checkEmail.Error.Error(),
		})

		c.Abort()
		return
	}

	//:: CHECK PASSWORD
	credentialError := user.CheckPassword(request.Password)

	if credentialError != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Password Not match",
			"error":   credentialError.Error(),
		})

		c.Abort()
		return
	}

	//:: GENERATE TOKEN
	tokenString, err := config.GenerateJWT(user.Email, user.Username, user.Role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token":   tokenString,
	})
}
