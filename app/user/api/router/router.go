package router

import (
	"gin_gomicro/app/user/api/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	api := r.Group("api")

	userGroup := api.Group("user")

	userGroup.POST("login", handlers.UserLogin)
	return r
}
