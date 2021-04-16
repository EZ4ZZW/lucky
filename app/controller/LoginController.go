package controller

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"lucky/app/common"
	"lucky/app/common/validate"
	"lucky/app/helper"
	"lucky/app/model"
	"net/http"
)

func DoLogin(c *gin.Context) {

	session := sessions.Default(c)

	if session.Get("idcard_number") != nil {
		data := make(map[string]interface{}, 0)
		_ = json.Unmarshal([]byte(session.Get("data").(string)), &data)
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "已登陆", data))
		return
	}

	var userModel = model.User{}
	var userValidate = validate.UserValidate

	var userJson model.User

	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据绑定模型错误", err.Error()))
		return
	}

	userMap := helper.Struct2Map(userJson)
	if res, err := userValidate.ValidateMap(userMap, "login"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "输入信息不完整或有误", err.Error()))
		return
	}

	res := userModel.LoginCheck(userJson)

	if res.Status == common.CodeSuccess {
		userInfo := res.Data.(model.User)
		returnData := map[string]interface{}{
			"user_id":       userInfo.ID,
			"user_name":     userInfo.Name,
			"idcard_number": userInfo.IdcardNumber,
			"user_school":   userInfo.School,
		}
		jsonData, _ := json.Marshal(returnData)
		session.Set("user_id", userInfo.ID)
		session.Set("user_name", userInfo.Name)
		session.Set("data", string(jsonData))
		session.Save()
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, returnData))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func DoLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "注销成功", session.Get("user_id")))
}
