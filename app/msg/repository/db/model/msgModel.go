package model

import "gorm.io/gorm"

type MsgModel struct {
	*gorm.Model
	UserId uint32 `gorm:"index; not null"`
	From   string `gorm:"not null"`
	Body   string `gorm:"not null"`
}
