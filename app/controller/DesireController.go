package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"lucky/app/common"
	"lucky/app/common/validate"
	"lucky/app/helper"
	"lucky/app/model"
	"net/http"
)

func AddDesire(c *gin.Context) {
	var desireJson model.Desire
	desireValidate := validate.DesireValidate
	desireModel := model.Desire{}
	userDesireModel := model.UserDesire{}
	var userDesireJson model.UserDesire

	if err := c.ShouldBindJSON(&desireJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	desireMap := helper.Struct2Map(desireJson)

	if res, err := desireValidate.ValidateMap(desireMap, "add"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	userDesireJson.UserID = int(helper.GetUserIdFromSession(c))
	res := desireModel.AddDesire(desireJson)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	userDesireJson.DesireID = res.Data.(model.Desire).ID
	res2 := userDesireModel.AddUserDesire(userDesireJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res2.Status, res2.Msg, res.Data))

}

func AchieveDesire(c *gin.Context) {
	var desireJson model.Desire
	desireModel := model.Desire{}
	desireValidate := validate.DesireValidate

	if err := c.ShouldBindJSON(&desireJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定失败", err.Error()))
		return
	}

	desireMap := helper.Struct2Map(desireJson)

	if res, err := desireValidate.ValidateMap(desireMap, "achieve"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	res := desireModel.AchieveDesire(desireJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func GetUserDesire(c *gin.Context) {
	var userDesireJson model.UserDesire
	desireModel := model.UserDesire{}
	userDesireValidate := validate.UserDesireValidate

	if err := c.ShouldBindJSON(&userDesireJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定失败", err.Error()))
		return
	}
	log.Print(userDesireJson)
	userDesireMap := helper.Struct2Map(userDesireJson)

	if res, err := userDesireValidate.ValidateMap(userDesireMap, "getUser"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	res := desireModel.GetUserAllDesire(userDesireJson)

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}
