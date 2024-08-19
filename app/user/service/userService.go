package service

import (
	"context"
	"errors"
	"gin_gomicro/app/user/model/do"
	"gin_gomicro/app/user/repository/db/dao"
	"gin_gomicro/idl/pb"
	"sync"

	"github.com/spf13/cast"
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
func (u *UserSrv) UserLogin(ctx context.Context, req *do.User) (uint64, error) {
	user, err := dao.NewUserDao(ctx).FindUserByAccount(req.Account)
	if err != nil {
		return 0, err
	}
	if user.Id == 0 {
		err = errors.New("用户不存在")
		return 0, err
	}
	if !user.CheckPassword(req.Password) {
		err = errors.New("密码错误")
		return 0, err
	}
	return user.Id, nil
}

// CheckPassword 密码验证
func (u *UserSrv) CheckPassword(ctx context.Context, req *do.User) (bool, error) {
	user, err := dao.NewUserDao(ctx).FindUserById(cast.ToUint64(req.Id))
	if err != nil {
		return false, err
	}
	ok := user.CheckPassword(user.Password)
	return ok, nil

}

// UserRegister 注册模块
func (u *UserSrv) UserRegister(ctx context.Context, req *do.User) (*do.User, error) {
	user, err := dao.NewUserDao(ctx).FindUserByAccount(req.Account)
	if err != nil {
		return nil, err
	}
	if user != nil {
		if user.Id > 0 {
			err := errors.New("账号已存在")
			return user, err
		}
	}

	user = &do.User{
		Account: req.Account,
	}
	if err := user.SetPassword(req.Password); err != nil {
		return nil, err
	}
	if err := dao.NewUserDao(ctx).UserRegister(user); err != nil {
		return nil, err
	}
	return user, nil
}

// GetAllUser 获取所有的user账户信息
func (u *UserSrv) GetAllUser(ctx context.Context, req *do.User) ([]*do.User, error) {
	userList, err := dao.NewUserDao(ctx).GetAllUser()
	if err != nil {
		return nil, err
	}
	return userList, nil
}

// 修改密码
func (u *UserSrv) ChangePassword(ctx context.Context, req *do.User) error {
	err := dao.NewUserDao(ctx).ChangePassword(req.Id, req.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserSrv) ChangeName(ctx context.Context, req *do.User) (*do.User, error) {
	user, err := dao.NewUserDao(ctx).FindUserById(req.Id)
	if err != nil {
		return nil, err
	}
	if !user.CheckPassword(req.Password) {
		return nil, errors.New("验证密码错误")
	}
	err = dao.NewUserDao(ctx).ChangeName(req.Id, req.Account)
	if err != nil {
		return nil, errors.New("更新失败")
	}
	user.Account = req.Account
	return user, nil
}
