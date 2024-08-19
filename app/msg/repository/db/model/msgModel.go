package model

import "gorm.io/gorm"

type Msg struct {
	*gorm.Model
	UserId uint32 `gorm:"index; not null"`
	From   uint32 `gorm:""`
	Body   string `gorm:"not null"`
	To     uint32 `gorm:""`
}
