package database

import (
	"github.com/mattot-the-builder/totblog/internal/model"
)

func CreatePost(post *model.Post) {
	DB.Create(post)
}

func GetAllPosts() []model.Post {
	var posts []model.Post
	DB.Find(&posts)
	return posts
}

func GetPostBySlug(slug string) model.Post {
	var post model.Post
	DB.Where("slug = ?", slug).First(&post)
	return post
}

func UpdatePost(slug string, post *model.Post) {
	DB.Model(&model.Post{}).Where("slug = ?", slug).Updates(post)
}

func DeletePost(slug string) {
	DB.Where("slug = ?", slug).Delete(&model.Post{})
}
