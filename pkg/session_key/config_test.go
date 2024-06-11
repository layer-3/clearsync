package session_key

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateKernelVersion(t *testing.T) {
	t.Run("Should accept supported versions", func(t *testing.T) {
		tcs := []string{
			"0.2.2",
			"0.2.4",
		}

		for _, version := range tcs {
			err := ValidateKernelVersion(version)
			assert.NoError(t, err)
		}
	})

	t.Run("Should reject unsupported versions", func(t *testing.T) {
		tcs := []string{
			"0.2.1",
			"0.3.0",
			"hello",
			"",
			"%!$@#$%",
		}

		for _, version := range tcs {
			err := ValidateKernelVersion(version)
			assert.Error(t, err)
		}
	})
}
