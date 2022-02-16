package vwap

import (
	"context"
	"vwap-calculator/coinbase"
)

type service struct {
	cbClient coinbase.CoinbaseClient
	pairs    []string
}

func NewService(cbClient coinbase.CoinbaseClient, pairs []string) Service {
	return &service{
		cbClient: cbClient,
		pairs:    pairs,
	}
}

func (s *service) Run(ctx context.Context) error {
	return nil
}
