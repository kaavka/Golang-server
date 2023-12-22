package main

import (
	"golang/server/constants"
	"golang/server/controller"
	"golang/server/service"

	"github.com/gin-gonic/gin"
)

var(
	userService service.UsersService = service.New()
	UserController controller.UserController = controller.New(userService)
)

func main() {
	server := gin.Default()

	server.GET(constants.USERS_ROUTE, UserController.FindAll)
	server.POST(constants.USERS_ROUTE, UserController.Save)
	server.PUT(constants.USERS_ROUTE, UserController.Edit)
	server.DELETE(constants.USERS_ROUTE, UserController.Delete)

	server.Run(constants.DEFAULT_PORT) 
}