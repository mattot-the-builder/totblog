package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Content string `json:"content"`
	User    User   `gorm:"foreignKey:UserID"`
}
