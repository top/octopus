package core

import (
	"octopus/core/data/parse"
	"octopus/core/data/pipeline"
	"octopus/core/data/resource"
	"octopus/core/data/scheduler"
	"octopus/core/global"
	"time"

	"go.uber.org/zap"
)

var Octo *Octopus

// Octopus main struct
type Octopus struct {
	taskname   string
	pScheduler scheduler.Scheduler
	mc         resource.Resource
	pParser    parse.Parse
	pPipelines []pipeline.Pipeline
}

// NewOctopus creates a new Octopus object
func NewOctopus(taskname string) *Octopus {
	return &Octopus{
		taskname:   taskname,
		pScheduler: scheduler.NewChan(),
		mc:         resource.NewChan(global.CONFIG.System.Transform.Capacity),
		pPipelines: make([]pipeline.Pipeline, 0),
	}
}

// GetTaskName implements global.task interface
func (o *Octopus) GetTaskName() string {
	return o.taskname
}

// SetScheduler sets the scheduler
func (o *Octopus) SetScheduler(s scheduler.Scheduler) *Octopus {
	o.pScheduler = s
	return o
}

// GetScheduler returns the scheduler
func (o *Octopus) GetScheduler() scheduler.Scheduler {
	return o.pScheduler
}

// GetPipelines returns the pipelines
func (o *Octopus) GetPipelines() []pipeline.Pipeline {
	return o.pPipelines
}

// AddPipelines adds new pipelines
func (o *Octopus) AddPipelines(pipeline ...pipeline.Pipeline) *Octopus {
	o.pPipelines = append(o.pPipelines, pipeline...)
	return o
}

// Close destructor
func (o *Octopus) Close() {
	if o.pScheduler != nil {
		o.pScheduler.Close()
	}
}

// Run execution
func (o *Octopus) Run() {
	for {
		pkg, err := o.pScheduler.Poll()
		if err == nil && pkg != nil {
			o.mc.GetOne()
			go func(p *scheduler.Package) {
				defer o.mc.FreeOne()
				o.process(p)
			}(pkg)
		}
		time.Sleep(1 * time.Second)
	}
}

// process implements the crawling
func (o *Octopus) process(pkg *scheduler.Package) {
	defer func() {
		if err := recover(); err != nil {
			global.LOGGER.Error("octopus process failed", zap.Any("err", err))
		}
	}()

	global.LOGGER.Info("octopus process with:", zap.Any("pkg", pkg))

	// initializing resource parser, currently supports URL and File, but can be extended by implementing Execute() method
	switch pkg.Source {
	case global.SOURCE_ARCHIVE:
		o.pParser = parse.Archive(pkg.Media)
	case global.SOURCE_API:
		fallthrough
	default:
		o.pParser = parse.App(pkg.Media)
	}

	// crawling and transforming data
	data, err := o.pParser.Execute(pkg)
	if err != nil {
		global.LOGGER.Error("octopus process parse execute failed", zap.Any("pkg", pkg), zap.Error(err))
		return
	}
	global.LOGGER.Debug("got data", zap.Any("data", data))

	// distributing to pipelines
	if data != nil {
		for _, p := range o.GetPipelines() {
			if err := p.Distribute(pkg, data); err != nil {
				global.LOGGER.Error("octopus pipeline distribute failed", zap.Any("pkg", pkg), zap.Any("data", data), zap.Error(err))
			}
		}
	}
}
