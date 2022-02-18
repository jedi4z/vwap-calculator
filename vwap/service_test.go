package vwap_test

import (
	"net/url"
	"strings"
	"testing"
	"vwap-calculator/coinbase"
	"vwap-calculator/vwap"

	"github.com/stretchr/testify/require"
)

func TestNewService(t *testing.T) {
	t.Parallel()

	e := url.URL{Scheme: "wss", Host: "ws-feed.exchange.coinbase.com", Path: "/"}
	pairsArr := strings.Split("BTC-USD,ETH-USD,ETH-BTC", ",")

	cbClient, _ := coinbase.NewClient(e.String())

	_, err := vwap.NewService(cbClient, pairsArr)

	require.NoError(t, err)
}
