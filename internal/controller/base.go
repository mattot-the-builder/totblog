package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mattot-the-builder/totblog/internal/database"
	"github.com/mattot-the-builder/totblog/internal/model"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (s *Server) Initialize() {

	// setup database

	s.DB = database.Connect()
	log.Println("database setup complete")

	// setup router

	s.Router = gin.Default()

	// setup routes

	SetupRoute(s.Router)

	log.Println("route setup complete")
}

func (s *Server) Run(addr string) {
	s.Router.Run(addr)
}

func (s *Server) Migrate() {
	s.DB.AutoMigrate(&model.User{})
	s.DB.AutoMigrate(&model.Post{})
}
