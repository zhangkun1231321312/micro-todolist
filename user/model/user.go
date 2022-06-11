package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PassWordDigest string
}

const (
	PasswordConst = 12 //密码加密难度
)

//加密密码
func (user *User) SetPassWord(password string) error {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), PasswordConst)
	if err != nil {
		return err
	}
	user.PassWordDigest = string(byte)
	return nil
}

//检验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWordDigest), []byte(password))
	return err == nil

}
