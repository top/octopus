package router

import (
	"octopus/core/service/api"

	"github.com/gin-gonic/gin"
)

func Manage(router *gin.RouterGroup) {
	archive := router.Group("manage")
	{
		archive.PUT("/:platform/auth", api.ManageAuth)
	}
}
