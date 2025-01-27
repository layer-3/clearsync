package driver

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/layer-3/clearsync/pkg/quotes/common"
)

func TestOnce_Start(t *testing.T) {
	t.Parallel()

	t.Run("Should call the function only once", func(t *testing.T) {
		t.Parallel()

		o := common.NewOnce()
		require.True(t, o.Start(func() {}))
		require.False(t, o.Start(func() {}), 1, "Start() method was executed more than once")
	})

	t.Run("Should reset the STOP action", func(t *testing.T) {
		t.Parallel()

		o := common.NewOnce()
		require.True(t, o.Start(func() {}))
		require.True(t, o.Stop(func() {}))
		require.True(t, o.Start(func() {}))
	})
}

func TestOnce_Stop(t *testing.T) {
	t.Parallel()

	t.Run("Should call the function only once", func(t *testing.T) {
		t.Parallel()

		o := common.NewOnce()
		require.True(t, o.Start(func() {})) // start the process to unblock STOP action
		require.True(t, o.Stop(func() {}))
		require.False(t, o.Stop(func() {}), "Stop() method was executed more than once")
	})

	t.Run("Should reset the START action", func(t *testing.T) {
		t.Parallel()

		o := common.NewOnce()
		stoppedChan := make(chan bool, 2)
		defer close(stoppedChan)

		require.True(t, o.Start(func() {})) // start the process to unblock STOP action
		require.True(t, o.Stop(func() {}))
		require.True(t, o.Start(func() {}))
		require.True(t, o.Stop(func() {}))
	})
}

func TestOnce_Subscribe(t *testing.T) {
	t.Parallel()

	t.Run("Should return false if Start has not been called", func(t *testing.T) {
		t.Parallel()

		o := common.NewOnce()
		require.False(t, o.Subscribe(), "Subscribe() should return false when Start() has not been called")
	})

	t.Run("Should return true if Start has been called", func(t *testing.T) {
		t.Parallel()

		o := common.NewOnce()
		require.True(t, o.Start(func() {}))
		require.True(t, o.Subscribe(), "Subscribe() should return true when Start() has been called")
	})

	t.Run("Should return false after Stop has been called", func(t *testing.T) {
		t.Parallel()

		o := common.NewOnce()
		require.True(t, o.Start(func() {}))
		require.True(t, o.Stop(func() {}))
		require.False(t, o.Subscribe(), "Subscribe() should return false after Stop() has been called")
	})
}

func TestOnce_Unsubscribe(t *testing.T) {
	t.Parallel()

	t.Run("Should return false if Start has not been called", func(t *testing.T) {
		t.Parallel()

		o := common.NewOnce()
		require.False(t, o.Unsubscribe(), "Unsubscribe() should return false when Start() has not been called")
	})

	t.Run("Should return true if Start has been called", func(t *testing.T) {
		t.Parallel()

		o := common.NewOnce()
		require.True(t, o.Start(func() {}))
		require.True(t, o.Unsubscribe(), "Unsubscribe() should return true when Start() has been called")
	})

	t.Run("Should return false after Stop has been called", func(t *testing.T) {
		t.Parallel()

		o := common.NewOnce()
		require.True(t, o.Start(func() {}))
		require.True(t, o.Stop(func() {}))
		require.False(t, o.Unsubscribe(), "Unsubscribe() should return false after Stop() has been called")
	})
}
