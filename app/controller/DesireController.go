package controller

import (
	"github.com/gin-gonic/gin"
	"lucky/app/common"
	"lucky/app/common/validate"
	"lucky/app/helper"
	"lucky/app/model"
	"net/http"
)

func AddDesire(c *gin.Context) {
	var desireJson model.Desire
	desireValite := validate.DesireValidate
	desireModel := model.Desire{}

	if err := c.ShouldBindJSON(&desireJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "绑定数据模型失败", err.Error()))
		return
	}

	desireMap := helper.Struct2Map(desireJson)

	if res,err := desireValite.ValidateMap(desireMap, "add"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	res := desireModel.AddDesire(desireJson)

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}

