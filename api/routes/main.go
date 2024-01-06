package routes

import (
	"crud/api/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	tweetController := controllers.NewTweetController()

	v1 := router.Group("")
	{
		v1.GET("/tweets", tweetController.FindAll)
		v1.POST("/tweet", tweetController.Create)
	}

	return v1

}
