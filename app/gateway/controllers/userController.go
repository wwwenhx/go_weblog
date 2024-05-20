package controllers

import (
	"context"
	"fmt"
	"gin_gomicro/app/gateway/grpc"
	"gin_gomicro/config"
	"gin_gomicro/idl/pb"
	"gin_gomicro/pkg/e"
	"gin_gomicro/pkg/etcd"
	"gin_gomicro/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

var rpc = grpc.GetRpc(etcd.GetServiceAddr("userService"))

// 用户登录
func (u *UserController) Login(c *gin.Context) {
	response := &utils.Response{}
	req := &pb.UserReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user, err := rpc.UserLogin(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseFail(err, "登录失败"))
	}
	token, err := utils.GenerateToken(user.UserDetail.Id)
	if err != nil {
		return
	}
	fmt.Println(token)
	dataMap := make(map[string]interface{}, 3)
	dataMap["token"] = token
	dataMap["userId"] = user.UserDetail.Id
	dataMap["userName"] = user.UserDetail.UserName
	c.JSON(http.StatusOK, response.ResponseSuccess(dataMap, "登录成功"))
}

// 校验密码
func (u *UserController) CheckPassword(c *gin.Context) {
	response := &utils.Response{}
	req := &pb.UserReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseFail(err.Error(), "请求失败"))
		return
	}
	fmt.Println(req)
	res, err := rpc.CheckPassword(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, response.ResponseFail(err, "校验密码失败"))
		return
	}
	fmt.Println(res)
	c.JSON(http.StatusOK, response.ResponseSuccess(res, "密码校验成功"))
}

// 获取全部user
func (u *UserController) GetAllUsers(c *gin.Context) {
	response := utils.Response{}
	rpc := grpc.GetRpc(config.UserServiceAddress)
	req := &pb.UserReq{}
	user, err := rpc.GetAllUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseFail(err, "不存在用户"))
		return
	}
	response.Data = user
	response.Code = e.Success
	c.JSON(http.StatusOK, response.ResponseSuccess(user, "获取数据成功"))
}

// 注册用户
func (u *UserController) RegisterUser(c *gin.Context) {
	response := utils.Response{}
	req := &pb.UserReq{}
	res := &pb.UserRes{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res, err := rpc.UserRegister(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	if res.Code != e.Success {
		c.JSON(http.StatusBadRequest, response.ResponseFail(err, ""))
	}
	response.Data = res
	response.Code = e.Success
	c.JSON(http.StatusOK, response.ResponseSuccess(res, "注册成功"))
}

// 修改密码
func (u *UserController) UpdatePassword(c *gin.Context) {
	response := utils.Response{}
	req := &pb.UserReq{}
	res := &pb.UserRes{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res, err := rpc.ChangePassword(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	if res.Code != e.Success {
		c.JSON(http.StatusBadRequest, response.ResponseFail(err, ""))
	}
	c.JSON(http.StatusOK, response.ResponseSuccess(res, "修改成功"))
}

// 修改账号
