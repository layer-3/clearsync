package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOnce_Start(t *testing.T) {
	t.Parallel()

	t.Run("Should call the function only once", func(t *testing.T) {
		t.Parallel()

		o := newOnce()
		startedChan := make(chan bool, 2)
		defer close(startedChan)

		o.Start(func() { startedChan <- true })
		o.Start(func() { startedChan <- true })

		require.Len(t, startedChan, 1, "Start() method was executed more than once")
	})

	t.Run("Should reset the STOP action", func(t *testing.T) {
		t.Parallel()

		o := newOnce()
		startedChan := make(chan bool, 2)
		defer close(startedChan)

		o.Start(func() { startedChan <- true })
		o.Stop(func() {})
		o.Start(func() { startedChan <- true })

		require.Len(t, startedChan, 2)
	})
}

func TestOnce_Stop(t *testing.T) {
	t.Parallel()

	t.Run("Should call the function only once", func(t *testing.T) {
		t.Parallel()

		o := newOnce()
		stoppedChan := make(chan bool, 2)
		defer close(stoppedChan)

		o.Start(func() {}) // start the process to unblock STOP action
		o.Stop(func() { stoppedChan <- true })
		o.Stop(func() { stoppedChan <- true })

		require.Len(t, stoppedChan, 1, "Stop() method was executed more than once")
	})

	t.Run("Should reset the START action", func(t *testing.T) {
		t.Parallel()

		o := newOnce()
		stoppedChan := make(chan bool, 2)
		defer close(stoppedChan)

		o.Start(func() {}) // start the process to unblock STOP action
		o.Stop(func() { stoppedChan <- true })
		o.Start(func() {})
		o.Stop(func() { stoppedChan <- true })

		require.Len(t, stoppedChan, 2)
	})
}
