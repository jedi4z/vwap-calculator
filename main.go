package main

import (
	"context"
	"log"
	"net/url"
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

	vwapService := vwap.NewService(cbClient)

	err = vwapService.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
