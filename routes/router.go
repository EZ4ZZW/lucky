package routes

import (
	"github.com/gin-gonic/gin"
	Controller "lucky/app/controller"
)

func Routes(router *gin.Engine) {

	api := router.Group("/")
	{
		api.POST("/Ccnulogin", Controller.CcnuLogin)
		api.POST("/login", Controller.DoLogin)
		api.POST("/user/email", Controller.BindEmail)

		wishes := api.Group("/wishes")
		{
			wishes.POST("", Controller.AddDesire)
			wishes.POST("/light", Controller.AchieveDesire)
			wishes.GET("", Controller.GetUserDesire)
			wishes.GET("/byID", Controller.GetWishByID)
			wishes.GET("/categories", Controller.GetWishByCatagories)
			wishes.DELETE("", Controller.DeleteWish)
		}

		message := api.Group("/message")
		{
			message.POST("/leave", Controller.LeaveMessage)
			message.GET("/get", Controller.GetUserMessage)
		}

		api.GET("")
	}

}
