package vwap

import (
	"context"
	"log"
	"vwap-calculator/coinbase"

	"golang.org/x/xerrors"
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
	receiver := make(chan coinbase.Response)

	err := s.cbClient.Subscribe(ctx, s.pairs, receiver)
	if err != nil {
		return xerrors.Errorf("service subscription err: %w", err)
	}

	for data := range receiver {
		log.Printf("data: %v", data)
	}

	return nil
}
