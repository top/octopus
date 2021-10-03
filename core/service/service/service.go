package service

import (
	"errors"
	"octopus/core"
	"octopus/core/data/scheduler"
	"octopus/core/global"
	"octopus/core/service/model"

	"go.uber.org/zap"
)

func service(platform, source, userID string, app *model.App, archive *model.Archive) error {
	if core.Octo == nil {
		return errors.New("instance has not been initialized")
	}
	global.LOGGER.Debug("preparing to push service",
		zap.String("platform", platform),
		zap.String("source", source),
		zap.String("userID", userID),
		zap.Any("app", app),
		zap.Any("archive", archive),
	)
	return core.Octo.GetScheduler().Push(&scheduler.Package{
		Media:   platform,
		Source:  source,
		UserID:  userID,
		App:     app,
		Archive: archive,
		Priority: func() uint {
			if app != nil && app.Priority != 0 {
				return app.Priority
			} else if archive != nil && archive.Priority != 0 {
				return archive.Priority
			}
			return 0
		}(),
	})
}
