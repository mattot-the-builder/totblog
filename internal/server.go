package internal

import "github.com/mattot-the-builder/totblog/internal/controller"

var server = controller.Server{}

func Run() {
	server.Initialize()
	server.Run(":8080")
}
