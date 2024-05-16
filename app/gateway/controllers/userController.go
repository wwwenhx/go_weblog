package controllers

import (
	"context"
	"fmt"
	"gin_gomicro/app/gateway/etcd"
	"gin_gomicro/app/gateway/grpc"
	"gin_gomicro/config"
	"gin_gomicro/idl/pb"
	"gin_gomicro/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

var rpc = grpc.GetRpc(etcd.GetServiceAddr("userService"))

func (u *UserController) GetAllUsers(c *gin.Context) {
	rpc := grpc.GetRpc(config.UserServiceAddress)
	req := &pb.UserReq{}
	user, err := rpc.GetAllUser(context.Background(), req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserController) RegisterUser(c *gin.Context) {
	req := &pb.UserReq{}
	res := &pb.UserRes{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(rpc)
	res, err := rpc.UserRegister(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	if res.Code != e.Success {
		c.JSON(http.StatusBadRequest, gin.H{"Error:": "注册失败"})
	}
	c.JSON(http.StatusOK, res)
}
