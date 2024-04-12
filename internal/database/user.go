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

func UpdateUser(user *model.User) {
	DB.Model(user).Updates(user)
}

func DeleteUser(user *model.User) {
	DB.Delete(user, user.ID)
}
