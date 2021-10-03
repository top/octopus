package resource

type Chan struct {
	capnum int
	mc     chan int
}

func NewChan(num int) *Chan {
	return &Chan{
		capnum: num,
		mc:     make(chan int, num),
	}
}

func (c *Chan) GetOne() {
	c.mc <- 1
}

func (c *Chan) FreeOne() {
	<-c.mc
}

func (c *Chan) Has() int {
	return len(c.mc)
}

func (c *Chan) Left() int {
	return c.capnum - c.Has()
}
