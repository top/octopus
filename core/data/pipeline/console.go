package pipeline

import (
	"octopus/core/data/scheduler"
	"octopus/core/global"

	"go.uber.org/zap"
)

type Console struct{}

func NewConsole() *Console {
	return &Console{}
}

func (c *Console) Distribute(p *scheduler.Package, r *Result) error {
	global.LOGGER.Info(global.CONFIG.System.Zap.Prefix, zap.Any("scheduler package", p), zap.Any("pipeline result", r))
	return nil
}
