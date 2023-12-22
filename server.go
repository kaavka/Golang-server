package main

import (
	"golang/server/constants"
	"golang/server/controller"
	"golang/server/service"
	"golang/server/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var(
	userService service.UsersService = service.New()
	UserController controller.UserController = controller.New(userService)
)

func main() {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	port := utils.EnvPortOr(constants.DEFAULT_PORT)
	address := constants.HOST + port

	server.GET(constants.USERS_ROUTE, UserController.FindAll)
	server.POST(constants.USERS_ROUTE, UserController.Save)
	server.PUT(constants.USERS_ROUTE, UserController.Edit)
	server.DELETE(constants.USERS_ROUTE, UserController.Delete)

	server.Run(address) 
}