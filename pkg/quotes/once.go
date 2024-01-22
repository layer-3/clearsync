package quotes

import (
	"sync"
	"sync/atomic"
)

// Once manages the synchronization of starting and stopping a process.
// It ensures that the start and stop functions are called IN ORDER and ONLY ONCE.
type once struct {
	start sync.Once
	stop  sync.Once

	// Using plain bool is not thread-safe in this case,
	// since sync.Once executes passed function in a synchronized way,
	// but Subscribe and Unsubscribe may be called from async context.
	started atomic.Bool
}

func newOnce() *once {
	o := &once{}
	o.Stop()
	return o
}

func (o *once) Start() bool {
	// The value is not loaded from atomic storage here,
	// since there may be subsequent calls to Start
	// but sync.Once guarantees that the passed function is executed only once,
	var started bool
	o.start.Do(func() {
		started = true
		o.started.Store(true)
		o.stop = sync.Once{} // allow a new stop
	})
	return started
}

func (o *once) Stop() bool {
	// The value is not loaded from atomic storage here,
	// since there may be subsequent calls to Stop
	// but sync.Once guarantees that the passed function is executed only once,
	var stopped bool
	o.stop.Do(func() {
		stopped = true
		o.started.Store(false)
		o.start = sync.Once{} // allow a new start
	})
	return stopped
}

// Subscribe checks if Start has been called before allowing subscription
func (o *once) Subscribe() bool {
	return o.started.Load()
}

// Unsubscribe checks if Start has been called before allowing unsubscription
func (o *once) Unsubscribe() bool {
	return o.started.Load()
}
