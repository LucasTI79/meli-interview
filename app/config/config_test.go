package config_test

import (
	"os"
	"testing"
	"time"

	"github.com/lucasti79/meli-interview/config"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig_Defaults(t *testing.T) {
	// Clear environment variables to test defaults
	os.Unsetenv("SERVER_HOST")
	os.Unsetenv("SERVER_PORT")
	os.Unsetenv("SERVER_TIMEOUT_READ")
	os.Unsetenv("SERVER_TIMEOUT_WRITE")
	os.Unsetenv("SERVER_TIMEOUT_IDLE")

	cfg := config.LoadConfig()

	require.Equal(t, "0.0.0.0", cfg.Server.Host)
	require.Equal(t, "8080", cfg.Server.Port)
	require.Equal(t, 5*time.Second, cfg.Server.TimeoutRead)
	require.Equal(t, 5*time.Second, cfg.Server.TimeoutWrite)
	require.Equal(t, 5*time.Second, cfg.Server.TimeoutIdle)
}

func TestLoadConfig_FromEnv(t *testing.T) {
	os.Setenv("SERVER_HOST", "127.0.0.1")
	os.Setenv("SERVER_PORT", "9000")
	os.Setenv("SERVER_TIMEOUT_READ", "10")
	os.Setenv("SERVER_TIMEOUT_WRITE", "15")
	os.Setenv("SERVER_TIMEOUT_IDLE", "20")

	defer func() {
		// Clean up environment variables after test
		os.Unsetenv("SERVER_HOST")
		os.Unsetenv("SERVER_PORT")
		os.Unsetenv("SERVER_TIMEOUT_READ")
		os.Unsetenv("SERVER_TIMEOUT_WRITE")
		os.Unsetenv("SERVER_TIMEOUT_IDLE")
	}()

	cfg := config.LoadConfig()

	require.Equal(t, "127.0.0.1", cfg.Server.Host)
	require.Equal(t, "9000", cfg.Server.Port)
	require.Equal(t, 10*time.Second, cfg.Server.TimeoutRead)
	require.Equal(t, 15*time.Second, cfg.Server.TimeoutWrite)
	require.Equal(t, 20*time.Second, cfg.Server.TimeoutIdle)
}
