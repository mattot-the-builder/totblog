package database

import (
	"github.com/mattot-the-builder/totblog/internal/model"
)

func CreateUser(user *model.User) {
	DB.Create(user)
}

func GetAllUsers() []model.User {
	var users []model.User
	DB.Find(&users)
	return users
}

func GetUserByUsername(username string) model.User {
	var user model.User
	DB.Where("username = ?", username).First(&user)
	return user
}

func UpdateUser(username string, user *model.User) {
	DB.Model(&model.User{}).Where("username = ?", username).Updates(user)
}

func DeleteUser(username string) {
	DB.Where("username = ?", username).Delete(&model.User{})
}
