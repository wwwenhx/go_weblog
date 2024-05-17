package routers

import (
	"gin_gomicro/app/gateway/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	userController := &controllers.UserController{}
	router.POST("/user/login", userController.Login)
	r := router.Group("user")
	r.Use(AuthMiddleware())
	r.POST("/register", userController.RegisterUser)
	r.POST("/getAllUsers", userController.GetAllUsers)
}
