package scheduler

import (
	"errors"
	"octopus/core/global"
)

type Chan struct {
	ch chan *Package
}

func NewChan() *Chan {
	return &Chan{
		ch: make(chan *Package, global.CONFIG.System.Transform.Capacity),
	}
}

func (c *Chan) Push(pkg *Package) error {
	if c.Count() >= cap(c.ch) {
		return errors.New("channel full")
	}
	c.ch <- pkg
	return nil
}

func (c *Chan) Poll() (*Package, error) {
	if c.Count() == 0 {
		return nil, errors.New("channel empty")
	}
	return <-c.ch, nil
}

func (c *Chan) Count() int {
	return len(c.ch)
}

func (c *Chan) Close() {
	c.ch = make(chan *Package)
}
