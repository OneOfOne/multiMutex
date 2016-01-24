package multiMutex

import (
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const DefaultCleanupDuration = time.Minute

type rwmux struct {
	sync.RWMutex
	cnt int64
}

type part struct {
	m map[string]*rwmux
	sync.Mutex
}

type RWMutex struct {
	p []part
	t *time.Ticker
}

func NewRW() *RWMutex {
	return NewRWSize(0, 0)
}

func NewRWSize(sz int, cleanupDur time.Duration) (rw *RWMutex) {
	if sz <= 0 {
		sz = (runtime.NumCPU() + 1) * 2
	}
	if cleanupDur <= 0 {
		cleanupDur = DefaultCleanupDuration
	}
	rw = &RWMutex{
		p: make([]part, sz),
		t: time.NewTicker(cleanupDur),
	}
	for i := range rw.p {
		rw.p[i].m = map[string]*rwmux{}
	}
	return
}

func (rw *RWMutex) get(path string) (v *rwmux) {
	p := rw.p[modDjb2(path)%len(rw.p)]
	p.Lock()
	if v = p.m[path]; v == nil {
		v = &rwmux{}
		p.m[path] = v
	}
	v.Add(1)
	p.Unlock()
	return
}

func (rw *RWMutex) Lock(path string) func() {
	mux := rw.get(path)
	mux.Lock()
	return func() { mux.Unlock(); mux.Add(-1) }
}

func (rw *RWMutex) RLock(path string) func() {
	mux := rw.get(path)
	mux.RLock()
	return func() { mux.RUnlock(); mux.Add(-1) }
}

func (rw *rwmux) Add(v int64) {
	atomic.AddInt64(&rw.cnt, v)
}
