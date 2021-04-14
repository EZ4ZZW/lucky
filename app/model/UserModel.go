package model

import (
	"lucky/app/common"
	"lucky/app/helper"
)

type User struct {
	ID           int    `json:"id" gorm:"id"`
	IdcardNumber string `json:"idcard_number" gorm:"idcard_number"`
	Password     string `json:"password" gorm:"password"`
	School       string `json:"school"`
	Wechat       string `json:"wechat"`
	Name         string `json:"name"`
	Tel          string `json:"tel"`
	Email        string `json:"email"`
	Major        string `json:"major"`
}

func (User) TableName() string {
	return "user"
}

func (model *User) LoginCheck(data User) helper.ReturnType {
	user := User{}
	err := db.Where("student_number = ? AND password = ?", data.IdcardNumber, data.Password).First(&user).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "用户名或密码错误", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "登录验证成功", Data: user}
	}
}
