package dao

import (
	"context"
	"fmt"
	"gin_gomicro/app/user/model/do"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	db := NewDBClient(ctx)
	if db == nil {
		fmt.Println("db client is nil")
	}
	return &UserDao{db}
}

// FindUserByName 通过名字查询用户
func (dao *UserDao) FindUserByAccount(name string) (user *do.User, err error) {
	err = dao.Model(&do.User{}).Where("user_name=?", name).First(&user).Error
	return
}

// FindUserById 通过id查询用户
func (dao *UserDao) FindUserById(id uint64) (user *do.User, err error) {
	err = dao.Model(&do.User{}).Where("id=?", id).First(&user).Error
	return
}

// 用户注册
func (dao *UserDao) UserRegister(user *do.User) (err error) {
	err = dao.Model(&do.User{}).Create(&user).Error
	return
}

// 获取所有user
func (dao *UserDao) GetAllUser() (userList []*do.User, err error) {
	err = dao.Model(&do.User{}).Find(&userList).Error
	return
}

// 修改密码
func (dao *UserDao) ChangePassword(id uint64, password string) (err error) {
	err = dao.Model(&do.User{}).Where("id = ?", id).Update("password", password).Error
	return
}

// 修改用户名
func (dao *UserDao) ChangeName(id uint64, name string) (err error) {
	err = dao.Model(&do.User{}).Where("id = ?", id).Update("name", name).Error
	return
}
