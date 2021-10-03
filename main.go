package main

import (
	"octopus/core"
	"octopus/core/data/pipeline"
	"octopus/core/data/scheduler"
	"octopus/core/global"
	"octopus/core/helper"
	"octopus/core/initial"

	"go.uber.org/zap"
)

func main() {
	// initializing configuration and logger
	if err := initial.Viper(); err != nil {
		global.LOGGER.Fatal("failed to load config file", zap.Error(err))
	}
	global.LOGGER = initial.Zap()
	global.LOGGER.Debug("config content:", zap.Any("config", global.CONFIG))

	// initializing Redis
	global.REDIS = helper.NewRedis(global.CONFIG.Redis.Host, global.CONFIG.Redis.Pass, global.CONFIG.Redis.DB)
	if global.REDIS == nil {
		global.LOGGER.Fatal("redis initialization failed")
	}

	// initializing app Facebook
	global.FACEBOOK = initial.Facebook(
		global.CONFIG.Media.Facebook.AppKey,
		global.CONFIG.Media.Facebook.AppSecret,
		global.CONFIG.Media.Facebook.RedirectURI,
	)
	if global.FACEBOOK == nil {
		global.LOGGER.Fatal("facebook initialization failed")
	}

	// starting web server
	go func() {
		global.LOGGER.Fatal("server start failed", zap.Error(initial.StartServer()))
	}()

	// creating main thread
	core.Octo = core.NewOctopus("octopus")
	if core.Octo == nil {
		global.LOGGER.Fatal("octopus initialization failed")
	}

	// gc on exit
	defer func() {
		global.LOGGER.Info("preparing to quit...")
		core.Octo.Close()
		println("bye ~")
	}()

	// starting app
	core.Octo.SetScheduler(
		scheduler.NewRedis(
			global.CONFIG.Redis.Host,
			global.CONFIG.Redis.Pass,
			global.CONFIG.Redis.DB,
			global.CONFIG.Redis.Scheduler,
		)).
		AddPipelines(pipeline.NewConsole()).
		Run()
}
