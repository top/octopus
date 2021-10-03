package scheduler

import (
	"errors"
	"octopus/core/global"
	"octopus/core/helper"

	"go.uber.org/zap"
)

type Redis struct {
	rc      *helper.Redis
	channel string
}

func NewRedis(addr, pass string, db int, ch string) *Redis {
	return &Redis{
		rc:      helper.NewRedis(addr, pass, db),
		channel: ch,
	}
}

func (r *Redis) Push(pkg *Package) error {
	if r.Count() >= global.CONFIG.System.Transform.Capacity {
		return errors.New("redis list is full")
	}
	if pkg != nil && pkg.UserID != "" { // TODO:
		r.rc.RPush(r.channel, pkg)
		return nil
	}
	return errors.New("redis list push failed")
}

func (r *Redis) Poll() (*Package, error) {
	if r.Count() > 0 {
		pkg := &Package{}
		err := helper.NewJson().Unmarshal([]byte(r.rc.Poll(r.channel)), pkg)
		if err != nil {
			return nil, err
		}
		return pkg, nil
	}
	return nil, errors.New("list empty")
}

func (r *Redis) Count() int {
	return r.rc.Count(r.channel)
}

func (r *Redis) Close() {
	if err := r.rc.Close(); err != nil {
		global.LOGGER.Error("redis close failed", zap.Error(err))
	}
}
