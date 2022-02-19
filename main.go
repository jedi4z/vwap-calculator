package main

import (
	"context"
	"flag"
	"log"
	"net/url"
	"strings"
	"vwap-calculator/coinbase"
	"vwap-calculator/vwap"
)

const (
	defaultInterval = 200
	defaultPairs    = "BTC-USD,ETH-USD,ETH-BTC"
)

func main() {
	ctx := context.Background()

	pairs := flag.String("pairs", defaultPairs, "trading pairs to subscribe to")
	interval := flag.Int("interval", defaultInterval, "window size")

	flag.Parse()

	pairsArr := strings.Split(*pairs, ",")

	e := url.URL{Scheme: "wss", Host: "ws-feed.exchange.coinbase.com", Path: "/"}

	cbClient, err := coinbase.NewClient(e.String())
	if err != nil {
		log.Fatal(err)
	}

	vwapService, err := vwap.NewService(cbClient, pairsArr, *interval)
	if err != nil {
		log.Fatal(err)
	}

	err = vwapService.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
