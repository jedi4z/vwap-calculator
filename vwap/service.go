package vwap

import (
	"context"
	"log"
	"strconv"
	"vwap-calculator/coinbase"

	"golang.org/x/xerrors"
)

type service struct {
	cbClient   coinbase.CoinbaseClient
	pairs      []string
	vwapPeriod VWAPPeriod
	interval   int
}

func NewService(cbClient coinbase.CoinbaseClient, pairs []string, interval int) (Service, error) {
	vwapPeriod, err := NewVWAPPeriod(interval)
	if err != nil {
		return &service{}, xerrors.Errorf("error creating the VWAP period: %w", err)
	}

	return &service{
		cbClient:   cbClient,
		pairs:      pairs,
		vwapPeriod: vwapPeriod,
		interval:   interval,
	}, nil
}

func generateDataPointFromCoinbaseResponse(d coinbase.Response) (DataPoint, error) {
	priceFloat, err := strconv.ParseFloat(d.Price, 64)
	if err != nil {
		return DataPoint{}, xerrors.Errorf("error converting str price to float %s: %w", d.Price, err)
	}

	volumeFloat, err := strconv.ParseFloat(d.Size, 64)
	if err != nil {
		return DataPoint{}, xerrors.Errorf("error converting str volume to float %s: %w", d.Price, err)
	}

	return DataPoint{
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

	log.Printf("collecting data points (%d per pair), the VWAP values will be displayed soon", s.interval)

	for data := range receiver {
		if data.ProductID == "" || data.Price == "" {
			continue
		}

		datapoint, err := generateDataPointFromCoinbaseResponse(data)
		if err != nil {
			return xerrors.Errorf("error generating a data point from coinbase response: %w", err)
		}

		s.vwapPeriod.Calculate(datapoint)

		if vwap := s.vwapPeriod.GetVWAP(); len(vwap) != 0 {
			log.Println(vwap)
		}
	}

	return nil
}
