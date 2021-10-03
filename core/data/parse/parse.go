package parse

import (
	"octopus/core/data/parse/app"
	"octopus/core/data/pipeline"
	"octopus/core/data/scheduler"
	"octopus/core/global"
)

type Parse interface {
	Execute(pkg *scheduler.Package) (*pipeline.Result, error)
}

func App(platform string) Parse {
	switch platform {
	case global.MEDIA_FACEBOOK:
		return app.NewFacebook()
	case global.MEDIA_TWITTER:
		return app.NewTwitter()
	case global.MEDIA_INSTAGRAM:
		return app.NewInstagram()
	}
	return nil
}

func Archive(platform string) Parse {
	return nil
}
