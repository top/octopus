package scheduler

import (
	"octopus/core/service/model"
)

type Package struct {
	Media    string // 平台
	Source   string // 来源
	UserID   string // 账户
	App      *model.App
	Archive  *model.Archive
	Priority uint // 优先级
}

type Scheduler interface {
	Push(pkg *Package) error
	Poll() (*Package, error)
	Count() int
	Close()
}
