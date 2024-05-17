package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	Password string `gorm:"not null"`
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
