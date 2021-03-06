package tango

import (
	"sync"
	"reflect"
)

type pool struct {
	size int
	tp reflect.Type
	pool reflect.Value
	cur int
	lock sync.Mutex
}

func newPool(size int, tp reflect.Type) *pool {
	return &pool{
		size: size,
		cur: 0,
		pool: reflect.MakeSlice(reflect.SliceOf(tp), size, size),
		tp: reflect.SliceOf(tp),
	}
}

func (p *pool) New() reflect.Value {
	p.lock.Lock()
	defer func() {
		p.cur++
		p.lock.Unlock()
	}()

	if p.cur == p.pool.Len() {
		p.pool = reflect.MakeSlice(p.tp, p.size, p.size)
		p.cur = 0
	}
	return p.pool.Index(p.cur).Addr()
}