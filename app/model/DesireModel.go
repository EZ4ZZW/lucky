package model

import (
	"github.com/gin-gonic/gin"
	"lucky/app/common"
	"lucky/app/helper"
	"time"
)

type Desire struct {
	ID            int       `json:"wish_id" gorm:"id"`
	Desire        string    `json:"wish" gorm:"desire"`
	LightAt       time.Time `json:"light_at" gorm:"light_at"`
	CreatAt       time.Time `json:"creat_at" gorm:"creat_at"`
	WishmanName   string    `json:"wishman_name" gorm:"wishman_qq"`
	WishmanQQ     string    `json:"wishman_qq" gorm:"wishman_qq"`
	WishmanWechat string    `json:"wishman_wechat" gorm:"wishman_wechat"`
	WishmanTel    string    `json:"wishman_tel" gorm:"wishman_tel"`
	State         int       `json:"state" gorm:"state"`
	Type          int       `json:"type" gorm:"type"`
}

func (Desire) TableName() string {
	return "desire"
}

func (model *Desire) AddDesire(data Desire) helper.ReturnType {

	err := db.Model(&Desire{}).Omit("creat_at", "light_at").Create(&data).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "添加愿望失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "添加愿望成功", Data: data}
}

func (model *Desire) AchieveDesire(data Desire) helper.ReturnType {

	var desire Desire

	err := db.Model(&Desire{}).Where("id = ?", data.ID).Find(&desire).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "实现愿望失败", Data: err.Error()}
	}

	if desire.State == 0 {
		desire.State = 1
		desire.LightAt = time.Now()
		err := db.Model(&Desire{}).Update(&desire).Error
		if err != nil {
			return helper.ReturnType{Status: common.CodeError, Msg: "实现愿望失败,数据库错误", Data: err.Error()}
		}
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "实现愿望成功", Data: gin.H{
			"wish_state": true,
		}}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "实现愿望失败，该愿望已经被别人抢先实现了", Data: 1}
}
