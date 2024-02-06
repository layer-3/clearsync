package safe

// https://github.com/tidwall/spinlock

import (
	"runtime"
	"sync"
	"sync/atomic"
)

// An RWSpinLock is a reader/writer mutual exclusion lock.
// The lock can be held by an arbitrary number of readers
// or a single writer.
// RWSpinLockes can be created as part of other
// structures; the zero value for a RWSpinLock is
// an unlocked mutex.
type RWSpinLock struct {
	_     sync.Mutex // for copy protection compiler warning
	state uint32
}

const (
	rwmutexUnlocked       = 0
	rwmutexWrite          = 1 << 0 // Bit 1 is used as a flag for write mode
	rwmutexReadOffset     = 1 << 1 // Bits 2-32 store the number of readers
	rwmutexUnderflow      = ^uint32(rwmutexWrite)
	rwmutexWriterUnset    = ^uint32(rwmutexWrite - 1)
	rwmutexReaderDecrease = ^uint32(rwmutexReadOffset - 1)
)

// RLock locks rw for reading.
func (rw *RWSpinLock) RLock() {
	// Increase the number of readers by 1
	state := atomic.AddUint32(&rw.state, rwmutexReadOffset)

	// If no write bits are set, the read lock was successfully acquired
	if state&rwmutexWrite == 0 {
		return
	}

	// Otherwise we have to wait until the write bits become unset.
	// Afterwards the RWSpinLock is in read mode.
	for {
		if state := atomic.LoadUint32(&rw.state); state&rwmutexWrite == 0 {
			return
		}
		runtime.Gosched()
	}
}

// TryRLock tries to lock rw for reading.
// If a lock for reading can not be acquired immediately, false is returned.
func (rw *RWSpinLock) TryRLock() bool {
	// Increase the number of readers by 1
	state := atomic.AddUint32(&rw.state, rwmutexReadOffset)

	// If no write bits are set, the read lock was successfully acquired
	if state&rwmutexWrite == 0 {
		return true
	}

	// Undo
	atomic.AddUint32(&rw.state, rwmutexReaderDecrease)
	return false
}

// RUnlock undoes a single RLock call;
// it does not affect other simultaneous readers.
// It is a run-time error if rw is not locked for reading
// on entry to RUnlock.
func (rw *RWSpinLock) RUnlock() {
	// Decrease the number of readers by 1
	state := atomic.AddUint32(&rw.state, rwmutexReaderDecrease)

	// Check for underflow
	if state&rwmutexUnderflow == rwmutexUnderflow {
		fatal("spinlock: RUnlock of unlocked RWSpinLock")
	}
}

// Lock locks rw for writing.
// If the lock is already locked for reading or writing,
// Lock blocks until the lock is available.
func (rw *RWSpinLock) Lock() {
	for !atomic.CompareAndSwapUint32(&rw.state, rwmutexUnlocked, rwmutexWrite) {
		runtime.Gosched()
	}
}

// TryLock tries to lock rw for writing.
// If the lock for writing can not be acquired immediately, false is returned.
func (rw *RWSpinLock) TryLock() bool {
	return atomic.CompareAndSwapUint32(&rw.state, rwmutexUnlocked, rwmutexWrite)
}

// Unlock unlocks rw for writing.  It is a run-time error if rw is
// not locked for writing on entry to Unlock.
//
// As with Mutexes, a locked RWSpinLock is not associated with a particular
// goroutine.  One goroutine may RLock (Lock) an RWSpinLock and then
// arrange for another goroutine to RUnlock (Unlock) it.
func (rw *RWSpinLock) Unlock() {
	// Unset the Write bit
	state := atomic.AddUint32(&rw.state, rwmutexWriterUnset)
	if state&rwmutexWrite > 0 {
		fatal("sync: Unlock of unlocked RWSpinLock")
	}
}

// RLocker returns a Locker interface that implements
// the Lock and Unlock methods by calling rw.RLock and rw.RUnlock.
func (rw *RWSpinLock) RLocker() sync.Locker {
	return (*rlocker)(rw)
}

type rlocker RWSpinLock

func (r *rlocker) Lock()   { (*RWSpinLock)(r).RLock() }
func (r *rlocker) Unlock() { (*RWSpinLock)(r).RUnlock() }
