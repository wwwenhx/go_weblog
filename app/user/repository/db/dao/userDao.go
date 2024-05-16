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
	fmt.Println("enter NewUserDao")
	if ctx == nil {
		ctx = context.Background()
	}
	db := NewDBClient(ctx)
	if db == nil {
		fmt.Println("db client is nil")
	}
	fmt.Println("db client is ok")
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) FindUserByUserName(userName string) (user *model.User, err error) {
	err = dao.Model(&model.User{}).Where("user_name=?", userName).First(&user).Error
	return
}

func (dao *UserDao) UserRegister(user *model.User) (err error) {
	err = dao.Model(&model.User{}).Create(&user).Error
	return
}

func (dao *UserDao) GetAllUser() (userList []*model.User, err error) {
	err = dao.Model(&model.User{}).Find(&userList).Error
	return
}
