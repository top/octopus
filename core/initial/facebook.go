package initial

import (
	"fmt"
	"octopus/core/global"
	"octopus/core/helper"
	"octopus/core/helper/facebook"
	"octopus/core/service/model"

	"go.uber.org/zap"
)

func Facebook(appKey, appSecret, redirectUri string) *facebook.Session {
	global.LOGGER.Info("initializing Facebook with appKey", zap.String("appKey", appKey))
	app := facebook.New(appKey, appSecret, redirectUri)

	accessToken, err := global.REDIS.Get(fmt.Sprintf("%s_auth", global.MEDIA_FACEBOOK))
	if err != nil {
		global.LOGGER.Error("global redis get failed", zap.Error(err))
		return nil
	}
	var auth model.Auth
	if err := helper.NewJson().Unmarshal([]byte(accessToken), &auth); err != nil {
		global.LOGGER.Error("failed to unmarshal facebook access token", zap.Error(err))
		return nil
	}
	global.LOGGER.Debug("got access token from redis", zap.String("accessToken", auth.AccessToken))
	return app.Session(auth.AccessToken)
}
