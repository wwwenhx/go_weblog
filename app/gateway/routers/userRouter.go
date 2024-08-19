package routers

import (
	"gin_gomicro/app/gateway/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	userController := &controllers.UserController{}
	router.POST("/user/login", userController.Login)
	router.POST("/user/register", userController.RegisterUser)
	r := router.Group("user")
	// r.Use(AuthMiddleware())
	r.POST("/getAllUsers", userController.GetAllUsers)
	r.POST("/updatePassword", userController.UpdatePassword)
	r.POST("/checkPassword", userController.CheckPassword)
}
