package handlers

import (
	"gin_gomicro/app/user/model/do"
	"gin_gomicro/app/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func UserLogin(c *gin.Context) {
	var user do.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	userId, err := service.GetUserSrv().UserLogin(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.String(http.StatusOK, cast.ToString(userId))
}

func UserRegister(c *gin.Context) {
	var user do.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	userInfo, err := service.GetUserSrv().UserRegister(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, userInfo)
}
