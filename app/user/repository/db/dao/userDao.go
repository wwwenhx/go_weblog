package dao

import (
	"context"
	"fmt"
	"gin_gomicro/app/user/repository/db/model"
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

// 通过名字查询用户
func (dao *UserDao) FindUserByName(name string) (user *model.User, err error) {
	err = dao.Model(&model.User{}).Where("user_name=?", name).First(&user).Error
	return
}

// 通过id查询用户
func (dao *UserDao) FindUserById(id uint32) (user *model.User, err error) {
	err = dao.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return
}

// 用户注册
func (dao *UserDao) UserRegister(user *model.User) (err error) {
	err = dao.Model(&model.User{}).Create(&user).Error
	return
}

// 获取所有user
func (dao *UserDao) GetAllUser() (userList []*model.User, err error) {
	err = dao.Model(&model.User{}).Find(&userList).Error
	return
}

// 修改密码
func (dao *UserDao) ChangePassword(id uint32, password string) (err error) {
	err = dao.Model(&model.User{}).Where("id = ?", id).Update("password", password).Error
	return
}

// 修改用户名
func (dao *UserDao) ChangeName(id uint32, name string) (err error) {
	err = dao.Model(&model.User{}).Where("id = ?", id).Update("name", name).Error
	return
}
