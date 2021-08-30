package controller

import (
	"encoding/json"
	"lucky/app/common"
	"lucky/app/common/validate"
	"lucky/app/helper"
	"lucky/app/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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

// 强制用户绑定邮箱
func BindEmail(c *gin.Context) {
	var userJson model.User
	userModel := model.User{}

	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "数据绑定失败", err.Error()))
	}

	token := c.Request.Header.Get("token")
	student_number, err := helper.VerifyToken(token)
	if err == nil {
		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "token解析失败", student_number))
	}

	userJson.IdcardNumber = student_number

	userMap := helper.Struct2Map(userJson)
	userValidator := validate.UserValidate

	if res, _ := userValidator.ValidateMap(userMap, "email"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", student_number))
		return
	}

	if res := userModel.BindEmail(userJson); res.Status == common.CodeError {
		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "邮箱绑定失败", err.Error()))
	}

}

func CcnuLogin(c *gin.Context) {
	var userJson model.User
	userModel := model.User{}

	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "数据绑定失败", err.Error()))
	}

	userMap := helper.Struct2Map(userJson)

	usreValidator := validate.UserValidate

	if boo, err := usreValidator.ValidateMap(userMap, "login"); !boo {
		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
	}

	// 首次登陆，验证一站式
	// 首次登陆
	if res := userModel.CcnuLogin(userJson); res.Status == common.CodeSuccess {
		// 生成token
		token := helper.CreatToken(userJson.IdcardNumber)
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "登陆成功", token))
	}

	if res := userModel.CcnuLogin(userJson); res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "登陆失败", nil))
	}

}
