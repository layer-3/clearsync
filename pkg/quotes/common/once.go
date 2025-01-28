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
	// Calling Stop upfront to reset the state.
	if err := o.Stop(func() error { return nil }); err != nil {
		panic(err) // this should never happen if implementation is correct
	}
	return o
}

// Start starts the process and calls the passed function.
// It returns true if the process was started successfully.
func (o *Once) Start(f func() error) error {
	// If error value won't be changed, then Do method was not called
	// therefore the process is already running.
	err := ErrAlreadyStarted
	o.start.Do(func() {
		err = f() // overriding the error value
		o.started.CompareAndSwap(false, true)
		o.stop = sync.Once{} // allow a new stop
	})
	return err
}

// Stop stops the process and calls the passed function.
// It returns true if the process was stopped successfully.
func (o *Once) Stop(f func() error) error {
	// If error value won't be changed, then Do method was not called
	// therefore the process is already running.
	err := ErrAlreadyStopped
	o.stop.Do(func() {
		err = f() // overriding the error value
		o.started.CompareAndSwap(true, false)
		o.start = sync.Once{} // allow a new start
	})
	return err
}

// IsStarted checks if the process has been started.
func (o *Once) IsStarted() bool {
	return o.started.Load()
}
