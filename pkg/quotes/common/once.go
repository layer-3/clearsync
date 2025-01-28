package common

import (
	"sync"
	"sync/atomic"
)

// Once manages the synchronization of starting and stopping a process.
// It ensures that:
// 1. methods are called IN ORDER: Start -> { Subscribe | Unsubscribe } -> Stop -> Start -> ...
// 2. the Start and Stop methods are executed ONLY ONCE.
type Once struct {
	start sync.Once
	stop  sync.Once

	// Using plain bool is not thread-safe in this case,
	// since sync.Once executes passed function in a synchronized way,
	// but Subscribe and Unsubscribe may be called from async context.
	started atomic.Bool
}

func NewOnce() *Once {
	o := &Once{}
	o.Stop(func() {}) // calling Stop to initialize the stop sync.Once
	return o
}

// Start starts the process and calls the passed function.
// It returns true if the process was started successfully.
func (o *Once) Start(f func()) bool {
	// The value is not loaded from atomic storage here,
	// since there may be subsequent calls to Start
	// but sync.Once guarantees that the passed function is executed only once,
	var started bool
	o.start.Do(func() {
		f()
		started = true
		o.started.Store(true)
		o.stop = sync.Once{} // allow a new stop
	})
	return started
}

// Stop stops the process and calls the passed function.
// It returns true if the process was stopped successfully.
func (o *Once) Stop(f func()) bool {
	// The value is not loaded from atomic storage here,
	// since there may be subsequent calls to Stop
	// but sync.Once guarantees that the passed function is executed only once,
	var stopped bool
	o.stop.Do(func() {
		f()
		stopped = true
		o.started.Store(false)
		o.start = sync.Once{} // allow a new start
	})
	return stopped
}

// Subscribe checks if Start has been called before allowing subscription.
func (o *Once) Subscribe() bool {
	return o.started.Load()
}

// Unsubscribe checks if Start has been called before allowing unsubscription.
func (o *Once) Unsubscribe() bool {
	return o.started.Load()
}
