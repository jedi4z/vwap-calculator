package main

import (
	"context"
	"log"
	"net/url"
	"strings"
	"vwap-calculator/coinbase"
	"vwap-calculator/vwap"
)

func main() {
	ctx := context.Background()

	e := url.URL{Scheme: "wss", Host: "ws-feed.exchange.coinbase.com", Path: "/"}

	cbClient, err := coinbase.NewClient(e.String())
	if err != nil {
		log.Fatal(err)
	}

	pairsArr := strings.Split("BTC-USD,ETH-USD,ETH-BTC", ",")
	vwapService := vwap.NewService(cbClient, pairsArr)

	err = vwapService.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
