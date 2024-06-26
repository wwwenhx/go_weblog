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
func (u *UserSrv) UserLogin(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	res := &pb.UserRes{}
	res.Code = e.Success
	user, err := dao.NewUserDao(ctx).FindUserByName(req.UserName)
	if err != nil {
		return res, err
	}
	if user.ID == 0 {
		err = errors.New("用户不存在")
		return res, err
	}
	if !user.CheckPassword(req.Password) {
		err = errors.New("密码错误")
		res.Code = e.Success
		return res, err
	}

	res.UserDetail = BuildUser(user)
	return res, nil
}

// CheckPassword 密码验证
func (u *UserSrv) CheckPassword(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	fmt.Println(req.Id)
	res := &pb.UserRes{}
	res.Code = e.Success
	user, err := dao.NewUserDao(ctx).FindUserById(req.Id)
	fmt.Println(user)
	if err != nil {
		return res, err
	}
	if !user.CheckPassword(req.Password) {
		res.Code = e.Error
		return res, errors.New("密码校验失败")
	}
	return res, nil
}

// UserRegister 注册模块
func (u *UserSrv) UserRegister(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	fmt.Println("enter UserRegister...")
	res := &pb.UserRes{}
	res.Code = e.Success

	userByName, _ := dao.NewUserDao(ctx).FindUserByName(req.UserName)
	if userByName != nil {
		if userByName.ID > 0 {
			res.Code = e.Error
			err := errors.New("账号已存在")
			return res, err
		}
	}

	user := &model.User{
		UserName: req.UserName,
	}
	if err := user.SetPassword(req.Password); err != nil {
		res.Code = e.Error
		return res, err
	}
	if err := dao.NewUserDao(ctx).UserRegister(user); err != nil {
		res.Code = e.Error
		return res, err
	}
	return res, nil
}

// GetAllUser 获取所有的user账户信息
func (u *UserSrv) GetAllUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	res := &pb.UserRes{}
	res.Code = e.Success
	userList, _ := dao.NewUserDao(ctx).GetAllUser()
	userModels := make([]*pb.UserModel, 0, len(userList))
	for _, user := range userList {
		userModels = append(userModels, BuildUser(user))
	}
	res.UserList = userModels
	return res, nil
}

// 修改密码
func (u *UserSrv) ChangePassword(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	res := &pb.UserRes{}
	res.Code = e.Success
	result := dao.NewUserDao(ctx).ChangePassword(req.Id, req.Password)
	if result.Error != nil {
		res.Code = e.Error
		return res, errors.New("更新失败")
	}
	return res, nil
}

func (u *UserSrv) ChangeName(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	res := &pb.UserRes{}
	res.Code = e.Success
	user, err := dao.NewUserDao(ctx).FindUserById(req.Id)
	if err != nil {
		return res, err
	}
	if !user.CheckPassword(req.Password) {
		res.Code = e.Error
		return res, errors.New("验证密码错误")
	}
	result := dao.NewUserDao(ctx).ChangeName(req.Id, req.UserName)
	if result.Error != nil {
		res.Code = e.Error
		return res, errors.New("更新失败")
	}
	user.UserName = req.UserName
	res.UserDetail = BuildUser(user)
	return res, nil
}

func BuildUser(u *model.User) *pb.UserModel {
	return &pb.UserModel{
		Id:        uint32(u.ID),
		UserName:  u.UserName,
		CreatedAt: u.CreatedAt.Unix(),
	}
}
