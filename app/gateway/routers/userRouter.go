package routers

import (
	"gin_gomicro/app/gateway/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	userController := &controllers.UserController{}
	r := router.Group("user")
	r.POST("", func(c *gin.Context) {})
	r.POST("/register", userController.RegisterUser)
}
