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
	desireJson := model.Desire{}
	desireValite := validate.DesireValidate
	desireModel := model.Desire{}

	if err := c.ShouldBindJSON(&desireJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(http.StatusOK, "绑定数据模型失败", err.Error()))
		return
	}

}