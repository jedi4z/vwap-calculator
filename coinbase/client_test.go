package coinbase_test

import (
	"net/url"
	"testing"
	"vwap-calculator/coinbase"

	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	t.Parallel()

	e := url.URL{Scheme: "wss", Host: "ws-feed.exchange.coinbase.com", Path: "/"}
	_, err := coinbase.NewClient(e.String())

	require.NoError(t, err)
}
