package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattot-the-builder/totblog/internal/auth"
	"github.com/mattot-the-builder/totblog/internal/database"
	"github.com/mattot-the-builder/totblog/internal/model"
)

func RegisterUser(c *gin.Context) {

	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	database.CreateUser(&user)

	// TODO: remove this after testing
	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"user":    user,
	})
}

func Login(c *gin.Context) {

	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	// authenticate user

	var result = database.GetUserByUsername(user.Username)

	if result.Username != user.Username || result.Password != user.Password {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid credentials",
		})
	}

	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "logged in",
		"token":   token,
	})
}
