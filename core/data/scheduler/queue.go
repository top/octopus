package scheduler

import (
	"container/list"
	"crypto/md5"
	"errors"
	"sync"
)

type Queue struct {
	lock  *sync.Mutex
	rm    bool
	rmKey map[[md5.Size]byte]*list.Element
	queue *list.List
}

func NewQueue(rmDuplicate bool) *Queue {
	return &Queue{
		lock:  &sync.Mutex{},
		rm:    rmDuplicate,
		rmKey: make(map[[md5.Size]byte]*list.Element),
		queue: list.New(),
	}
}

func (q *Queue) Push(pkg *Package) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	var key [md5.Size]byte
	if q.rm {
		key = md5.Sum([]byte(pkg.UserID)) // TODO: 以谁为key
		if _, ok := q.rmKey[key]; ok {
			return nil
		}
	}
	e := q.queue.PushBack(pkg)
	if q.rm {
		q.rmKey[key] = e
	}
	return nil
}

func (q *Queue) Poll() (*Package, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.Count() == 0 {
		return nil, errors.New("queue empty")
	}
	e := q.queue.Front()
	pkg := e.Value.(*Package)
	key := md5.Sum([]byte(pkg.UserID)) // TODO:
	q.queue.Remove(e)
	if q.rm {
		delete(q.rmKey, key)
	}
	return pkg, nil
}

func (q *Queue) Count() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.Len()
}

func (q *Queue) Close() {
	q = &Queue{
		lock:  &sync.Mutex{},
		rm:    true,
		rmKey: make(map[[md5.Size]byte]*list.Element),
		queue: list.New(),
	}
}
