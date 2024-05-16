package service

import (
	"context"
	"errors"
	"fmt"
	"gin_gomicro/app/user/repository/db/dao"
	"gin_gomicro/app/user/repository/db/model"
	"gin_gomicro/idl/pb"
	"gin_gomicro/pkg/e"
	"sync"
)

type UserSrv struct {
	pb.UnimplementedUserServer
}

var UserSrvIns *UserSrv
var UserSrvOne sync.Once

/*
懒汉式单例模式
*/

func GetUserSrv() *UserSrv {
	UserSrvOne.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

// Login 登录模块
func (u *UserSrv) UserLogin(ctx context.Context, req *pb.UserReq) (res *pb.UserRes, err error) {
	res.Code = e.Success
	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		return
	}
	if user.ID == 0 {
		err = errors.New("用户不存在")
		return
	}
	if !user.CheckPassword(req.Password) {
		err = errors.New("密码错误")
		res.Code = e.Success
		return
	}

	res.UserDetail = BuildUser(user)
	return
}

// UserRegister 注册模块
func (u *UserSrv) UserRegister(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	fmt.Println("enter UserRegister...")
	res := &pb.UserRes{}
	res.Code = e.Success

	userByName, _ := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if userByName != nil {
		if userByName.ID > 0 {
			res.Code = e.Error
			err := errors.New("账号已存在")
			return res, err
		}
	}

	user := &model.User{
		Username: req.UserName,
	}
	if err := userByName.SetPassword(userByName.Password); err != nil {
		res.Code = e.Error
		return res, err
	}
	user.Password = userByName.Password

	if err := dao.NewUserDao(ctx).UserRegister(user); err != nil {
		res.Code = e.Error
		return res, err
	}
	return res, nil
}

// GetAllUser 获取所有的user账户信息
func (u *UserSrv) GetAllUser(ctx context.Context, req *pb.UserReq) (res *pb.UserRes, err error) {
	res.Code = e.Success
	userList, _ := dao.NewUserDao(ctx).GetAllUser()
	userModels := make([]*pb.UserModel, 0, len(userList))
	for _, user := range userList {
		userModels = append(userModels, BuildUser(user))
	}
	res.UserList = userModels
	return
}

func BuildUser(u *model.User) *pb.UserModel {
	return &pb.UserModel{
		Id:        uint32(u.ID),
		UserName:  u.Username,
		CreatedAt: u.CreatedAt.Unix(),
	}
}
