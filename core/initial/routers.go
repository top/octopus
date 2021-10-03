package initial

import (
	"octopus/core/service/middleware"
	"octopus/core/service/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var r = gin.Default()
	r.Use(middleware.Activity()) // middleware.Cors()

	publicGroup := r.Group("octopus")
	{
		router.AppRouter(publicGroup)     // API应用接口
		router.ArchiveRouter(publicGroup) // 归档文件解析接口
		router.Manage(publicGroup)        // 管理接口
	}
	return r
}
