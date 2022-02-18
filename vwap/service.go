package vwap

import (
	"context"
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

func generateDataPointFromCoinbaseResponse(d coinbase.Response) (dataPoint, error) {
	priceFloat, err := strconv.ParseFloat(d.Price, 64)
	if err != nil {
		return dataPoint{}, xerrors.Errorf("error converting str price to float %s: %w", d.Price, err)
	}

	volumeFloat, err := strconv.ParseFloat(d.Size, 64)
	if err != nil {
		return dataPoint{}, xerrors.Errorf("error converting str volume to float %s: %w", d.Price, err)
	}

	return dataPoint{
		Pair:   d.ProductID,
		Price:  priceFloat,
		Volume: volumeFloat,
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

		datapoint, err := generateDataPointFromCoinbaseResponse(data)
		if err != nil {
			return xerrors.Errorf("error generating a data point from coinbase response: %w", err)
		}

		s.vwapPeriod.Calculate(datapoint)

		// fmt.Println(s.vwapPeriod.GetVWAP())
	}

	return nil
}
