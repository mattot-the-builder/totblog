package database

import (
	"github.com/mattot-the-builder/totblog/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := "mattot:password@tcp(127.0.0.1:3306)/totblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	Migrate()
}

func Migrate() {
	DB.AutoMigrate(&model.User{})
}
