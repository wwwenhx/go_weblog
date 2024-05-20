package dao

import (
	"context"
	"fmt"
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
