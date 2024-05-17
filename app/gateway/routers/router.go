package routers

import (
	"fmt"
	"gin_gomicro/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

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

func AuthMiddleware() gin.HandlerFunc {
	response := utils.Response{}
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			fmt.Println("token is empty")
			c.JSON(http.StatusUnauthorized, response.ResponseFail("err", "token校验失败"))
			c.Abort()
			return
		}
		tokenParts := strings.Split(token, " ")
		fmt.Println(tokenParts)
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, response.ResponseFail("err", "token校验失败"))
			c.Abort()
			return
		}
		claims, err := utils.VerifyToken(tokenParts[1])
		if claims == nil {
			c.JSON(http.StatusUnauthorized, response.ResponseFail(err, "token校验失败"))
			c.Abort()
			return
		}

		c.Set("userId", claims.UserId)
		c.Next()
	}
}
