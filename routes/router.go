package routes

import (
	"github.com/gin-gonic/gin"
	Controller "lucky/app/controller"
)

func Routes(router *gin.Engine) {

	api := router.Group("/")
	{
		api.POST("/login", Controller.DoLogin)
		api.POST("/user/email")

		wishes := api.Group("/wishes")
		{
			wishes.POST("", Controller.AddDesire)
			wishes.POST("/light", Controller.AchieveDesire)
			wishes.GET("", Controller.GetUserDesire)
			wishes.GET("/byID", Controller.GetWishByID)
			wishes.GET("/categories", Controller.GetWishByCatagories)
			wishes.DELETE("", Controller.DeleteWish)
		}

		api.GET("")
	}

}
