package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattot-the-builder/totblog/internal/database"
	"github.com/mattot-the-builder/totblog/internal/model"
)

func CreateUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	database.CreateUser(&user)
	c.JSON(http.StatusCreated, user)
}

func GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user := database.GetUserByUsername(username)
	c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
	users := database.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	username := c.Param("username")
	var user model.User
	c.BindJSON(&user)
	database.UpdateUser(username, &user)
	c.JSON(http.StatusOK, gin.H{})
}

func DeleteUser(c *gin.Context) {
	username := c.Param("username")
	database.DeleteUser(username)
	c.JSON(http.StatusOK, gin.H{})
}
