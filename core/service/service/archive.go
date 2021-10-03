package service

import (
	"octopus/core/global"
	"octopus/core/service/model"
)

func ArchiveTwitter(userID string, a model.Archive) error {
	return service(global.MEDIA_TWITTER, global.SOURCE_ARCHIVE, userID, nil, &a)
}

func ArchiveFacebook(userID string, a model.Archive) error {
	return service(global.MEDIA_FACEBOOK, global.SOURCE_ARCHIVE, userID, nil, &a)
}

func ArchiveInstagram(userID string, a model.Archive) error {
	return service(global.MEDIA_INSTAGRAM, global.SOURCE_ARCHIVE, userID, nil, &a)
}
