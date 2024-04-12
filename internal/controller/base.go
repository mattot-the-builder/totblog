package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (s *Server) Initialize() {

	// setup database

	dsn := "mattot:password@tcp(127.0.0.1:3306)/totblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	s.DB = db

	log.Println("connected to database")

	// setup router

	s.Router = gin.Default()

	// setup routes

	SetupRoute(s.Router)

	log.Println("route setup complete")
}

func (s *Server) Run(addr string) {
	s.Router.Run(addr)
}
