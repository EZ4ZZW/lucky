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

// 通过点击一个愿望查看 该愿望 详情
func GetWishByID(c *gin.Context) {
	// 定义一个 结构体 用来接收愿望详情
	var desireModel model.Desire

	// 定义一个 结构体 用来接 json格式 的愿望ID
	var desireJson model.Desire

	// 初始化一个 验证器 用来校验数据格式
	desireValidate := validate.DesireValidate

	if err := c.ShouldBindJSON(&desireJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定失败", err.Error()))
		return
	}

	// 将json数据的结构体转化为map，转化为map的目的只是为了校验格式
	desireMap := helper.Struct2Map(desireJson)

	if res, err := desireValidate.ValidateMap(desireMap, "achieve"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	res := desireModel.GetWishByID(desireJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}

// 通过点击分类查看 同种愿望
func GetWishByCatagories(c *gin.Context) {
	// 定义一个 结构体 用来接收愿望详情
	var desireModel model.Desire

	// 定义一个 结构体 用来接 json格式 的愿望type
	var desireJson model.Desire

	// 初始化一个 验证器 用来校验数据格式
	// desireValidate := validate.DesireValidate

	if err := c.ShouldBindJSON(&desireJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定失败", err.Error()))
		return
	}

	// 将json数据的结构体转化为map，转化为map的目的只是为了校验格式
	// desireMap := helper.Struct2Map(desireJson)

	// if res, err := desireValidate.ValidateMap(desireMap, "achieve"); !res {
	// 	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
	// 	return
	// }

	res := desireModel.GetWishByCategories(desireJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}

// 删除愿望
func DeleteWish(c *gin.Context) {
	// 定义一个 结构体 用来接收愿望详情
	var desireModel model.Desire

	// 定义一个 结构体 用来接 json格式 的愿望Type
	var desireJson model.Desire

	// 初始化一个 验证器 用来校验数据格式
	desireValidate := validate.DesireValidate

	if err := c.ShouldBindJSON(&desireJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定失败", err.Error()))
		return
	}

	// 将json数据的结构体转化为map，转化为map的目的只是为了校验格式
	desireMap := helper.Struct2Map(desireJson)

	if res, err := desireValidate.ValidateMap(desireMap, "achieve"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	res := desireModel.DeleteWish(desireJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}
