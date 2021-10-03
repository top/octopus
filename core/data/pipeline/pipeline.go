package pipeline

import "octopus/core/data/scheduler"

type Result struct {
	Profile interface{} `json:"profile"`
	Posts   interface{} `json:"posts"`
}

type Pipeline interface {
	Distribute(p *scheduler.Package, r *Result) error
}
