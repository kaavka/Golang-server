package main

import (
	"golang/server/constants"
	"golang/server/controller"
	"golang/server/service"
	"golang/server/utils"

	"github.com/gin-gonic/gin"
)

var(
	userService service.UsersService = service.New()
	UserController controller.UserController = controller.New(userService)
)

func main() {
	server := gin.Default()
	port := utils.EnvPortOr(constants.DEFAULT_PORT)
	address := constants.HOST + port

	server.GET(constants.USERS_ROUTE, UserController.FindAll)
	server.POST(constants.USERS_ROUTE, UserController.Save)
	server.PUT(constants.USERS_ROUTE, UserController.Edit)
	server.DELETE(constants.USERS_ROUTE, UserController.Delete)

	server.Run(address) 
}