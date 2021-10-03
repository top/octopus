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

func App(c *gin.Context) {
	var a model.App
	if err := c.ShouldBindQuery(&a); err != nil {
		global.LOGGER.Error("parameters are invalid", zap.Any("parameters", a), zap.Error(err))
		response.FailMessage(global.ERROR_INVALID_PARAM, "parameters are invalid", c)
		return
	}
	global.LOGGER.Debug("got app", zap.Any("app", a))

	if err := verify.Verify(a, verify.AppVerification); err != nil || helper.IsStructureEmpty(a, model.App{}) {
		global.LOGGER.Error("parameters are invalid", zap.Any("parameters", a), zap.Error(err))
		response.FailMessage(global.ERROR_INVALID_PARAM, "parameters are not verified", c)
		return
	}

	var (
		platform = c.Param("platform")
		userid  = c.Param("userid")
		err      error
	)
	global.LOGGER.Debug("got parameters from path", zap.Any("platform", platform), zap.Any("userid", userid))
	switch platform {
	case global.MEDIA_TWITTER:
		err = service.AppTwitter(userid, a)
	case global.MEDIA_FACEBOOK:
		err = service.AppFacebook(userid, a)
	case global.MEDIA_INSTAGRAM:
		err = service.AppInstagram(userid, a)
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
