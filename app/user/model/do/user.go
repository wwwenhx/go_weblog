package do

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id         uint64    `json:"id" gorm:"column:id;not null"`
	Account    string    `json:"account" gorm:"column:account;not null"`
	Password   string    `json:"password" gorm:"column:password; not null"`
	Gender     uint8     `json:"gender" gorm:"column:gender;not null"`
	Nick       string    `json:"nick" gorm:"column:nick;not null"`
	Birthday   time.Time `json:"birthday" gorm:"column:birthday;not null"`
	Deleted    uint      `json:"deleted" gorm:"column:deleted;not null"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;not null"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;not null"`
}

func (User) TableName() string {
	return "user"
}

const PasswordCost = 12

func (u *User) SetPassword(password string) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return
	}
	u.Password = string(bytes)
	err1 := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err1 != nil {
		fmt.Println("bcrypt error")
	}

	return
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
