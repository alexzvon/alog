package logger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	logger := New()

	require.NotNil(t, logger)

	err := logger.Info("Test Info Message")
	require.Nil(t, err)

	err = logger.Warn("Test Warn Message")
	require.Nil(t, err)

	err = logger.Error("Test Error Message")
	require.Nil(t, err)
}
