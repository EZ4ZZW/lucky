package model

import (
	"lucky/app/common"
	"lucky/app/helper"
	"time"
)

type Desire struct {
	ID int `json:"id" gorm:"id"`
	Desire string `json:"wish" gorm:"desire"`
	WishmanName string `json:"wish_man_name" gorm:"wishman_qq"`
	WishManWechat string `json:"wish_man_wechat" gorm:"wishman_wechat"`
	WishManTel string `json:"wish_man_tel" gorm:"wishman_tel"`
	State int `json:"state" gorm:"state"`
	CreatAt time.Time `json:"creat_at" gorm:"creat_at"`
	LightAt time.Time `json:"light_at" gorm:"light_at"`
}

func (Desire) TableName() string {
	return "desire"
}

func (model *Desire) AddDesire(data Desire) helper.ReturnType {

	err := db.Create(&data).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "添加愿望失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "添加愿望成功", Data: data}
}