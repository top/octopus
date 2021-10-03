package api

import (
	"errors"
	"octopus/core/global"
	"octopus/core/helper"
	"octopus/core/service/model"
	"octopus/core/service/model/response"
	"octopus/core/service/model/verify"
	"octopus/core/service/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Archive(c *gin.Context) {
	var a model.Archive
	if err := c.ShouldBindQuery(&a); err != nil {
		global.LOGGER.Error("parameters are invalid", zap.Error(err))
		response.FailMessage(global.ERROR_INVALID_PARAM, "parameters are invalid", c)
		return
	}

	if err := verify.Verify(a, verify.ArchiveVerification); err != nil || helper.IsStructureEmpty(a, model.Archive{}) {
		global.LOGGER.Error("parameters are invalid", zap.Error(err))
		response.FailMessage(global.ERROR_INVALID_PARAM, "parameters are not verified", c)
		return
	}

	var (
		platform = c.Param("platform")
		userid  = c.Param("userid")
		err      error
	)
	switch platform {
	case global.MEDIA_TWITTER:
		err = service.ArchiveTwitter(userid, a)
	case global.MEDIA_FACEBOOK:
		err = service.ArchiveFacebook(userid, a)
	case global.MEDIA_INSTAGRAM:
		err = service.ArchiveInstagram(userid, a)
	default:
		err = errors.New(global.ERROR_UNSUPPORTED_PLATFORM)
	}
	if err != nil {
		global.LOGGER.Error("service failed", zap.Error(err))
		response.FailMessage(global.ERROR_INTERNAL, err.Error(), c)
	} else {
		response.OK(c)
	}
}
