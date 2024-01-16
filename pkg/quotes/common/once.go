package common

import (
	"sync"
)

// Once manages the synchronization of starting and stopping a process.
// It ensures that the start and stop functions are called IN ORDER and ONLY ONCE.
type Once struct {
	start sync.Once
	stop  sync.Once
}

func NewOnce() *Once {
	o := &Once{}
	o.Stop(func() {})
	return o
}

func (o *Once) Start(f func()) {
	o.start.Do(func() {
		f()
		o.stop = sync.Once{} // allow a new stop
	})
}

func (o *Once) Stop(f func()) {
	o.stop.Do(func() {
		f()
		o.start = sync.Once{} // allow a new start
	})
}
