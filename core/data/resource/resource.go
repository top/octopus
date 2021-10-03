package resource

type Resource interface {
	GetOne()
	FreeOne()
	Has() int
	Left() int
}
