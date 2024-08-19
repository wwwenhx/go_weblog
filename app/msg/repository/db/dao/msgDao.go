package dao

import (
	"context"
	"fmt"
	"gin_gomicro/app/msg/repository/db/model"
	"gorm.io/gorm"
)

type MsgDao struct {
	*gorm.DB
}

func NewMsgDao(ctx context.Context) *MsgDao {
	if ctx == nil {
		ctx = context.Background()
	}
	db := NewDBClient(ctx)
	if db == nil {
		fmt.Println("db client is nil")
	}
	return &MsgDao{db}
}

func (dao *MsgDao) SendMsg(msg model.Msg) (err error) {
	err = dao.Model(&model.Msg{}).Create(&msg).Error
	return err
}
