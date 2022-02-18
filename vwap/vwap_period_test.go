package vwap_test

import (
	"testing"
	"vwap-calculator/vwap"

	"github.com/stretchr/testify/require"
)

func TestNewVWAPPeriod(t *testing.T) {
	t.Parallel()

	_, err := vwap.NewVWAPPeriod(200)

	require.NoError(t, err)
}
