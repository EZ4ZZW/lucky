package model

import (
	"github.com/gin-gonic/gin"
	"lucky/app/common"
	"lucky/app/helper"
)

type UserDesire struct {
	ID       int `json:"id" gorm:"id"`
	UserID   int `json:"user_id" gorm:"user_id"`
	DesireID int `json:"desire_id" gorm:"desire_id"`
}

func (UserDesire) TableName() string {
	return "user_desire"
}

func (model *UserDesire) AddUserDesire(data UserDesire) helper.ReturnType {

	err := db.Model(&UserDesire{}).Create(&data).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "添加用户愿望失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "添加用户愿望成功", Data: 0}

}

func (model *UserDesire) GetUserAllDesire(data UserDesire) helper.ReturnType {

	var userDesires []UserDesire

	err := db.Model(&UserDesire{}).Where("user_id = ?", data.UserID).Find(&userDesires).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询用户愿望失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "获取用户愿望成功", Data: gin.H{
		"wishes": userDesires,
	}}

}
