package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/mattot-the-builder/totblog/internal/database"
	"github.com/mattot-the-builder/totblog/internal/model"
)

func CreatePost(c *gin.Context) {
	var post model.Post
	c.BindJSON(&post)

	// create slug for the post title

	post.Slug = slug.Make(post.Title + " " + string(post.ID))

	database.CreatePost(&post)
	c.JSON(http.StatusCreated, post)
}

func GetPostBySlug(c *gin.Context) {
	postSlug := c.Param("slug")
	post := database.GetPostBySlug(postSlug)
	c.JSON(http.StatusOK, post)
}

func GetAllPosts(c *gin.Context) {
	posts := database.GetAllPosts()
	c.JSON(http.StatusOK, posts)
}

func UpdatePost(c *gin.Context) {
	postSlug := c.Param("slug")
	var post model.Post
	c.BindJSON(&post)
	database.UpdatePost(postSlug, &post)
	c.JSON(http.StatusOK, gin.H{})
}

func DeletePost(c *gin.Context) {
	slug := c.Param("slug")
	var post model.Post
	c.BindJSON(&post)
	database.DeletePost(slug)
	c.JSON(http.StatusOK, gin.H{})
}
