package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {
	r.POST("/login", Login)
	r.POST("/register", RegisterUser)
}
