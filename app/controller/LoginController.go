package controller

import (
	"log"
	"lucky/app/common"
	"lucky/app/common/validate"
	"lucky/app/helper"
	"lucky/app/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CallbackInfo struct {
	Sexname      string `form:"sexname"`
	School       int    `form:"school"`
	Deptcodename string `form:"deptcodename"`
	Cardno       string `form:"cardno"`
	Name         string `form:"name"`
	Sno          string `form:"sno"`
	Uclassname   string `form:"uclassname"`
	Zydm         string `form:"zydm"`
	Email        string `form:"email"`
	Zydmname     string `form:"zydmname"`
}

func Callback(c *gin.Context) {
	whutUserInfo := struct {
		UserInfo    CallbackInfo `form:"user"`
		Continueurl string       `form:"continueurl"`
	}{}

	userModel := model.User{}
	userJson := model.User{}
	userJson.School = 1

	if err := c.ShouldBind(&whutUserInfo); err != nil {
		c.JSON(http.StatusForbidden, helper.ApiReturn(common.CodeError, "回调失败", nil))
		return
	}
	log.Println(whutUserInfo)
	userJson.School = 1
	userJson.IdcardNumber = whutUserInfo.UserInfo.Sno
	userJson.Name = whutUserInfo.UserInfo.Name
	if whutUserInfo.UserInfo.Sexname == "男" {
		userJson.Gender = 0
	} else {
		userJson.Gender = 1
	}
	if whutUserInfo.UserInfo.Deptcodename != "" {
		userJson.Major = whutUserInfo.UserInfo.Deptcodename
	}

	res := userModel.GetUserByRealName(userJson.Name)

	// 首次登陆，新建用户
	if res.Status == common.CodeError {
		res = userModel.CreateUser(userJson)
		token := helper.CreatToken(strconv.Itoa(res.Data.(model.User).ID))
		c.SetCookie("jwt_token", token, 604800, "/", "localhost", false, false)
		c.Request.Header.Add("jwt_token", token)
		if whutUserInfo.Continueurl != "" {
			c.Redirect(http.StatusPermanentRedirect, whutUserInfo.Continueurl)
		} else {
			// 根据性别进行重定向
			if userJson.Gender == 0 {
				c.Redirect(http.StatusPermanentRedirect, "https://ipandai.club")
			} else {
				c.Redirect(http.StatusPermanentRedirect, "https://ipandai.club")
			}
		}
		return
	}

	// set http only as false
	// expire at 7 days before
	token := helper.CreatToken(strconv.Itoa(res.Data.(model.User).ID))
	c.SetCookie("jwt_token", token, 604800, "/", "localhost", false, false)
	c.Request.Header.Add("jwt_token", token)
	if whutUserInfo.Continueurl != "" {
		c.Redirect(http.StatusMovedPermanently, whutUserInfo.Continueurl)
		return
	}
	// TODO: 最后进行redirect
	//if(empty($_POST['continueurl'])) $this->redirect($data['gender'] == 0? 'https://lucky-day.itoken.team/act/login/boy_lead_page.html': 'https://lucky-day.itoken.team/act/login/girl_lead_page.html');
	//else $this->redirect($_POST['continueurl']);

}

func WhutLogin(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "http://ias.sso.itoken.team/portal.php?posturl=https%3A%2F%2Fipandai.club%2Fapi%2Flogin%2Fwhut%2Fcallback&continueurl=https://ipandai.club")
}

// 强制用户绑定邮箱
func BindEmail(c *gin.Context) {
	var userJson model.User
	userModel := model.User{}

	student_number := c.MustGet("student_number").(string)

	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "数据绑定失败", err.Error()))
		return
	}

	userJson.IdcardNumber = student_number

	userMap := helper.Struct2Map(userJson)
	userValidator := validate.UserValidate

	if res, err := userValidator.ValidateMap(userMap, "email"); !res {
		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	if res := userModel.BindEmail(userJson); res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "邮箱绑定失败", res.Data))
		return
	}

	c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeSuccess, "邮箱绑定成功", userJson.Email))

}

// ccnulogin
func CcnuLogin(c *gin.Context) {
	var userJson model.User
	userModel := model.User{}

	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "数据绑定失败", err.Error()))
		return
	}

	userMap := helper.Struct2Map(userJson)

	usreValidator := validate.UserValidate

	if boo, err := usreValidator.ValidateMap(userMap, "login"); !boo {
		c.JSON(http.StatusNotFound, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	// 首次登陆，验证一站式
	// 首次登陆
	if res := userModel.CcnuLogin(userJson); res.Status == common.CodeSuccess {
		// 生成token
		token := helper.CreatToken(userJson.IdcardNumber)
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "登陆成功", token))
	} else {
		c.JSON(http.StatusNotFound, helper.ApiReturn(res.Status, res.Msg, res.Data))
	}

}

// 返回用户信息
func ReturnUserInfo(c *gin.Context) {
	var user model.User
	id := c.MustGet("user_id").(int)
	res := user.GetInfoByUserID(id)
	if res.Status == common.CodeError {
		c.JSON(http.StatusNotFound, helper.ApiReturn(res.Status, res.Msg, res.Data))
	}

	c.JSON(http.StatusNotFound, helper.ApiReturn(res.Status, res.Msg, res.Data))
}
