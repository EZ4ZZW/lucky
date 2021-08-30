package model

import (
	"lucky/app/common"
	"lucky/app/helper"
)

type User struct {
	ID           int    `json:"id" gorm:"id"`
	IdcardNumber string `json:"idcard_number" gorm:"column:student_number"`
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

func (model *User) BindEmail(data User) helper.ReturnType {

	err := db.Model(&User{}).Where("student_number = ?", data.IdcardNumber).Update(&data).Error

	if err != nil {
		return helper.ReturnRes(common.CodeError, "绑定邮箱失败", err.Error())
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "绑定邮箱成功", Data: data.Email}
	}

}

//  CCNU Login登陆验证
func (model *User) CcnuLogin(data User) helper.ReturnType {
	user := User{}

	if err := db.Where("student_number = ?", data.IdcardNumber).First(&user); err != nil {
		_, err2 := GetUserInfoFormOne(data.IdcardNumber, data.Password)
		if err2 != nil {
			return helper.ReturnRes(common.CodeError, "用户名或密码错误", err.Error)
		}

		if err := db.Model(&User{}).Create(&data).Error; err != nil {
			return helper.ReturnRes(common.CodeError, "添加用户失败", err.Error)
		}

	} else if data.Password != user.Password {
		return helper.ReturnRes(common.CodeError, "用户名或密码错误", err.Error)
	}

	return helper.ReturnRes(common.CodeSuccess, "登陆成功", nil)

}
