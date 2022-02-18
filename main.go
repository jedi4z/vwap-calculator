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
	pairsArr := strings.Split("BTC-USD,ETH-USD,ETH-BTC", ",")

	e := url.URL{Scheme: "wss", Host: "ws-feed.exchange.coinbase.com", Path: "/"}

	cbClient, err := coinbase.NewClient(e.String())
	if err != nil {
		log.Fatal(err)
	}

	vwapService, err := vwap.NewService(cbClient, pairsArr)
	if err != nil {
		log.Fatal(err)
	}

	err = vwapService.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
