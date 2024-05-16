package routers

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	UserRouter(ginRouter)
	return ginRouter
}
