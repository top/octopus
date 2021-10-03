package api

import (
	"fmt"
	"octopus/core/global"
	"octopus/core/helper"
	"octopus/core/service/model"
	"octopus/core/service/model/response"
	"octopus/core/service/model/verify"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ManageAuth(c *gin.Context) {
	var a model.Auth
	if err := c.ShouldBindJSON(&a); err != nil {
		global.LOGGER.Error("parameters are invalid", zap.Any("parameters", a), zap.Error(err))
		response.FailMessage(global.ERROR_INVALID_PARAM, "parameters are invalid", c)
		return
	}

	if err := verify.Verify(a, verify.AuthVerification); err != nil || helper.IsStructureEmpty(a, model.Auth{}) {
		global.LOGGER.Error("parameters are invalid", zap.Any("parameters", a), zap.Error(err))
		response.FailMessage(global.ERROR_INVALID_PARAM, "parameters are not verified", c)
		return
	}

	var (
		platform = c.Param("platform")
		err      error
	)

	if platform == "" {
		global.LOGGER.Error("parameter platform is required", zap.Error(err))
		response.FailMessage(global.ERROR_INVALID_PARAM, "parameter platform is required", c)
		return
	}

	a.Media = platform
	if err := global.REDIS.SetExp(fmt.Sprintf("%s_auth", platform), a, -1); err != nil {
		global.LOGGER.Error("set auth failed", zap.Any("auth", a), zap.Error(err))
		response.FailMessage(global.ERROR_INTERNAL, global.ERROR_INTERNAL_MESSAGE, c)
		return
	}

	response.OK(c)
}
