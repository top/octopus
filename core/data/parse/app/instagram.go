package app

import (
	"octopus/core/data/pipeline"
	"octopus/core/data/scheduler"
)

type Instagram struct {
}

func NewInstagram() *Instagram {
	return &Instagram{}
}

func (i *Instagram) Execute(pkg *scheduler.Package) (*pipeline.Result, error) {
	return nil, nil
}
