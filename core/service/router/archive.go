package router

import (
	"octopus/core/service/api"

	"github.com/gin-gonic/gin"
)

func ArchiveRouter(router *gin.RouterGroup) {
	archive := router.Group("archive")
	{
		archive.GET("/:platform/:userid", api.Archive)
	}
}
