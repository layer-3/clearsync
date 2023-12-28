package quotes

import (
	"sync"
)

// once manages the synchronization of starting and stopping a process.
// It ensures that the start and stop functions are called IN ORDER and ONLY ONCE.
type once struct {
	startOnce sync.Once
	stopOnce  sync.Once
}

func newOnce() *once {
	sm := &once{}
	sm.Stop(func() {})
	return sm
}

func (sm *once) Start(f func()) {
	sm.startOnce.Do(func() {
		f()
		sm.stopOnce = sync.Once{} // allow a new stop
	})
}

func (sm *once) Stop(f func()) {
	sm.stopOnce.Do(func() {
		f()
		sm.startOnce = sync.Once{} // allow a new start
	})
}
