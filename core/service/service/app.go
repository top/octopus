package service

import (
	"octopus/core/global"
	"octopus/core/service/model"
)

func AppTwitter(userID string, i model.App) error {
	return service(global.MEDIA_TWITTER, global.SOURCE_API, userID, &i, nil)
}

func AppFacebook(userID string, i model.App) error {
	return service(global.MEDIA_FACEBOOK, global.SOURCE_API, userID, &i, nil)
}

func AppInstagram(userID string, i model.App) error {
	return service(global.MEDIA_INSTAGRAM, global.SOURCE_API, userID, &i, nil)
}
