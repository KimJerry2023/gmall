package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string `gorm:"unique"`
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string `gorm:"size:1000"`
	Money          string
}

const (
	PasswordCost        = 12 // 密码加密难度
	Active       string = "active"
)

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

func (user *User) AvatarURL() string {
	signedGetURL := user.Avatar
	return signedGetURL
}
