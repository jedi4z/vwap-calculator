package vwap

import (
	"context"
	"vwap-calculator/coinbase"
)

type service struct {
	cbClient coinbase.CoinbaseWsClient
}

func NewService(cbClient coinbase.CoinbaseWsClient) Service {
	return &service{
		cbClient: cbClient,
	}
}

func (s *service) Run(ctx context.Context) error {
	return nil
}
