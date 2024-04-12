package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {

	// auth
	r.POST("/login", Login)
	r.POST("/register", RegisterUser)

	// user
	r.GET("/users", GetAllUsers)
	r.GET("/users/:username", GetUserByUsername)
	r.POST("/users", CreateUser)
	r.PUT("/users/:username", UpdateUser)
	r.DELETE("/users/:username", DeleteUser)

	// post
	r.GET("/posts", GetAllPosts)
	r.GET("/posts/:slug", GetPostBySlug)
	r.POST("/posts", CreatePost)
	r.PUT("/posts/:slug", UpdatePost)
	r.DELETE("/posts/:slug", DeletePost)
}
