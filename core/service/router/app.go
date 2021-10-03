package router

import (
	"octopus/core/service/api"

	"github.com/gin-gonic/gin"
)

func AppRouter(router *gin.RouterGroup) {
	app := router.Group("app")
	{
		app.GET("/:platform/:userid", api.App)
	}
}
