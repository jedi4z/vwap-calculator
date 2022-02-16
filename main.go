package main

import (
	"log"
	"net/url"
	"vwap-calculator/coinbase"
)

func main() {
	e := url.URL{Scheme: "wss", Host: "ws-feed.exchange.coinbase.com", Path: "/"}

	_, err := coinbase.NewClient(e.String())
	if err != nil {
		log.Fatal(err)
	}
}
