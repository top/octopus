package app

import (
	"octopus/core/data/pipeline"
	"octopus/core/data/scheduler"
)

type Twitter struct {
}

func NewTwitter() *Twitter {
	return &Twitter{}
}

func (t *Twitter) Execute(pkg *scheduler.Package) (*pipeline.Result, error) {
	return nil, nil
}
