package models

import (
	"github.com/jinzhu/gorm"
	ec "github.com/shunjiecloud/account/errcode"
	"github.com/shunjiecloud/account/modules"
	"github.com/shunjiecloud/errors"
	"go.uber.org/zap"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"index"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
	Mail     string `json:"mail" gorm:"index"`
	Phone    string `json:"phone"`
}

//  创建用户
func AddUser(user *User) error {
	err := CanUserBeRegisted(user.Name, user.Mail)
	if err != nil {
		return errors.Adapt(err)
	}
	if err := modules.ModuleContext.DefaultDB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

//  邮箱是否存在
func IsEmailExisted(email string) (isExisted bool, err error) {
	count := 0
	err = modules.ModuleContext.DefaultDB.Model(User{}).Where("mail = ?", email).Count(&count).Error
	if err != nil {
		return false, errors.Adapt(err)
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

//  用户名是否存在
func IsUserNameExisted(name string) (isExisted bool, err error) {
	count := 0
	err = modules.ModuleContext.DefaultDB.Model(User{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		return false, errors.Adapt(err)
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

//  用户是否可注册
func CanUserBeRegisted(username, email string) error {
	isExisted, err := IsEmailExisted(email)
	if err != nil {
		return errors.Adapt(err)
	}
	if isExisted == true {
		return errors.New(ec.ErrEmailAlreadyUsed, zap.String("email", email))
	}
	isExisted, err = IsUserNameExisted(username)
	if err != nil {
		return errors.Adapt(err)
	}
	if isExisted == true {
		return errors.New(ec.ErrUserNameAlreadyUsed, zap.String("username", email))
	}
	return nil
}
