package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func nothing() error { return nil }

func TestOnce_Start(t *testing.T) {
	t.Parallel()

	t.Run("Should call the function only once", func(t *testing.T) {
		t.Parallel()

		o := NewOnce()
		require.NoError(t, o.Start(nothing))
		require.Error(t, o.Start(nothing), 1, "Start() method was executed more than once")
	})

	t.Run("Should reset the STOP action", func(t *testing.T) {
		t.Parallel()

		o := NewOnce()
		require.NoError(t, o.Start(nothing))
		require.NoError(t, o.Stop(nothing))
		require.NoError(t, o.Start(nothing))
	})
}

func TestOnce_Stop(t *testing.T) {
	t.Parallel()

	t.Run("Should call the function only once", func(t *testing.T) {
		t.Parallel()

		o := NewOnce()
		require.NoError(t, o.Start(nothing)) // start the process to unblock STOP action
		require.NoError(t, o.Stop(nothing))
		require.Error(t, o.Stop(nothing), "Stop() method was executed more than once")
	})

	t.Run("Should reset the START action", func(t *testing.T) {
		t.Parallel()

		o := NewOnce()
		stoppedChan := make(chan bool, 2)
		defer close(stoppedChan)

		require.NoError(t, o.Start(nothing)) // start the process to unblock STOP action
		require.NoError(t, o.Stop(nothing))
		require.NoError(t, o.Start(nothing))
		require.NoError(t, o.Stop(nothing))
	})
}

func TestOnce_Subscribe(t *testing.T) {
	t.Parallel()

	t.Run("Should return false if Start has not been called", func(t *testing.T) {
		t.Parallel()

		o := NewOnce()
		require.False(t, o.IsStarted(), "Should return false when Start() has not been called")
	})

	t.Run("Should return true if Start has been called", func(t *testing.T) {
		t.Parallel()

		o := NewOnce()
		require.NoError(t, o.Start(nothing))
		require.True(t, o.IsStarted(), "Should return true when Start() has been called")
	})

	t.Run("Should return false after Stop has been called", func(t *testing.T) {
		t.Parallel()

		o := NewOnce()
		require.NoError(t, o.Start(nothing))
		require.NoError(t, o.Stop(nothing))
		require.False(t, o.IsStarted(), "Should return false after Stop() has been called")
	})
}
