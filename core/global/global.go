package global

import (
	"octopus/core/config"
	"octopus/core/helper"
	"octopus/core/helper/facebook"

	"go.uber.org/zap"
)

var (
	CONFIG   config.Config
	LOGGER   *zap.Logger
	REDIS    *helper.Redis
	FACEBOOK *facebook.Session
)
