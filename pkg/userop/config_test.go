package userop

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewClientConfigFromEnv(t *testing.T) {
	t.Run("Should NOT override provided values", func(t *testing.T) {
		// Arrange
		t.Setenv("USEROP_CLIENT_PAYMASTER_CONFIG_TYPE", "pimlico_erc20")
		t.Setenv("USEROP_CLIENT_SMART_WALLET_TYPE", "kernel")

		defaultConfig := ClientConfig{}
		defaultConfig.Init()

		// Act
		config, err := NewClientConfigFromEnv()

		// Assert
		require.NoError(t, err)
		require.Equal(t, "pimlico_erc20", config.Paymaster.Type.String(),
			"values that are set in the environment should override the provided values")
		require.Equal(t, "kernel", config.SmartWallet.Type.String(),
			"values that are set in the environment should override the provided values")
		require.Equal(t, defaultConfig.Gas.MaxFeePerGasMultiplier, config.Gas.MaxFeePerGasMultiplier,
			"values that are not set in the environment should be set to default values")
	})
}
