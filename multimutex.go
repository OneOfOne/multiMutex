package multiMutex

import (
	"runtime"
	"sync"
)

type MultiMutex struct {
	ms []sync.RWMutex
}

func New() *MultiMutex {
	return NewSize(runtime.NumCPU() + 1)
}

func NewSize(sz int) *MultiMutex {
	return &MultiMutex{
		ms: make([]sync.RWMutex, sz),
	}
}

func (mm *MultiMutex) Get(key string) *sync.RWMutex {
	return &mm.ms[modDjb2(key)%len(mm.ms)]
}

func (mm *MultiMutex) Lock(key string) (unlock func()) {
	m := mm.Get(key)
	m.Lock()
	return m.Unlock
}

func (mm *MultiMutex) RLock(key string) (unlock func()) {
	m := mm.Get(key)
	m.Lock()
	return m.Unlock
}
