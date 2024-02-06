package safe

// https://github.com/tidwall/spinlock

import (
	"runtime"
	"sync"
	"sync/atomic"
)

func fatal(s string) {
	panic(s)
}

const (
	mutexUnlocked = 0
	mutexLocked   = 1
)

// A SpinLock is a mutual exclusion lock.
// SpinLockes can be created as part of other structures;
// the zero value for a SpinLock is an unlocked mutex.
type SpinLock struct {
	_     sync.Mutex // for copy protection compiler warning
	state int32
}

// Lock locks m.
// If the lock is already in use, the calling goroutine repetitively tries to
// acquire the lock until it is available (busy waiting).
func (m *SpinLock) Lock() {
	for !atomic.CompareAndSwapInt32(&m.state, mutexUnlocked, mutexLocked) {
		runtime.Gosched()
	}
}

// TryLock tries to lock m.
// If the lock is already in use, the lock is not acquired and false is
// returned.
func (m *SpinLock) TryLock() bool {
	return atomic.CompareAndSwapInt32(&m.state, mutexUnlocked, mutexLocked)
}

// Unlock unlocks m.
// It is a run-time error if m is not locked on entry to Unlock.
//
// A locked SpinLock is not associated with a particular goroutine.
// It is allowed for one goroutine to lock a SpinLock and then
// arrange for another goroutine to unlock it.
func (m *SpinLock) Unlock() {
	state := atomic.AddInt32(&m.state, -mutexLocked)
	if state != mutexUnlocked {
		fatal("spinlock: unlock of unlocked mutex")
	}
}
