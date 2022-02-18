package vwap

import (
	"context"
	"fmt"
	"strconv"
	"vwap-calculator/coinbase"

	"golang.org/x/xerrors"
)

type service struct {
	cbClient   coinbase.CoinbaseClient
	pairs      []string
	vwapPeriod VWAPPeriod
}

func NewService(cbClient coinbase.CoinbaseClient, pairs []string) (Service, error) {
	vwapPeriod, err := NewVWAPPeriod(200)
	if err != nil {
		return &service{}, xerrors.Errorf("error creating the VWAP period: %w", err)
	}

	return &service{
		cbClient:   cbClient,
		pairs:      pairs,
		vwapPeriod: vwapPeriod,
	}, nil
}

func (s *service) Run(ctx context.Context) error {
	receiver := make(chan coinbase.Response)

	err := s.cbClient.Subscribe(ctx, s.pairs, receiver)
	if err != nil {
		return xerrors.Errorf("service subscription err: %w", err)
	}

	for data := range receiver {
		if data.ProductID == "" || data.Price == "" {
			continue
		}

		priceFloat, err := strconv.ParseFloat(data.Price, 64)
		if err != nil {
			return xerrors.Errorf("error converting str price to float %s: %w", data.Price, err)
		}

		volumeFloat, err := strconv.ParseFloat(data.Size, 64)
		if err != nil {
			return xerrors.Errorf("error converting str volume to float %s: %w", data.Price, err)
		}

		s.vwapPeriod.Calculate(dataPoint{
			Pair:   data.ProductID,
			Price:  priceFloat,
			Volume: volumeFloat,
		})

		fmt.Println(s.vwapPeriod.GetVWAP())
	}

	return nil
}
