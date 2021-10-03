package app

import (
	"octopus/core/data/pipeline"
	"octopus/core/data/scheduler"
)

type Facebook struct {
}

func NewFacebook() *Facebook {
	return &Facebook{}
}

func (f *Facebook) Execute(pkg *scheduler.Package) (*pipeline.Result, error) {
	return nil, nil
}
