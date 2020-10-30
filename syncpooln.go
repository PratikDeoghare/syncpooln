package syncpooln

import (
	"sync"
)

type Pooln struct {
	new func(n int) interface{}
	l   sync.RWMutex
	m   map[int]*sync.Pool
}

// New
func New(newFunc func(int) interface{}) *Pooln {
	return &Pooln{
		m:   make(map[int]*sync.Pool),
		new: newFunc,
	}
}

// Get calls the `Get` method on the pool at `n` and returns the result.
func (p *Pooln) Get(n int) interface{} {
	x := p.pool(n).Get()
	return x
}

// Put adds `x` to the pool at `n` by calling the pool's `Put` method.
func (p *Pooln) Put(n int, x interface{}) {
	p.pool(n).Put(x)
}

// pool returns the pool at `n`. If there is no pool at `n`, it creates new one,
// puts it in the map and returns it.
func (p *Pooln) pool(n int) *sync.Pool {
	p.l.RLock()
	v, ok := p.m[n]
	p.l.RUnlock()

	if ok {
		return v
	}

	pn := &sync.Pool{
		New: func() interface{} {
			return p.new(n)
		},
	}

	p.l.Lock()
	p.m[n] = pn
	p.l.Unlock()
	return pn
}
