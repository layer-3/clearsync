package quotes

import (
	"sync"
)

// once manages the synchronization of starting and stopping a process.
// It ensures that the start and stop functions are called IN ORDER and ONLY ONCE.
type once struct {
	start sync.Once
	stop  sync.Once
}

func newOnce() *once {
	o := &once{}
	o.Stop(func() {})
	return o
}

func (o *once) Start(f func()) {
	o.start.Do(func() {
		f()
		o.stop = sync.Once{} // allow a new stop
	})
}

func (o *once) Stop(f func()) {
	o.stop.Do(func() {
		f()
		o.start = sync.Once{} // allow a new start
	})
}
