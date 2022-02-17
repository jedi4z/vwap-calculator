package vwap

import "context"

type DataPoint struct {
	Price     float64
	Volume    float64
	ProductID string
}

type Period struct {
	Interval   uint
	DataPoints []DataPoint
	SumPrice   map[string]float64
	SumVolume  map[string]float64
	VWAP       map[string]float64
}

type Service interface {
	Calculate(DataPoint)
	Run(context.Context) error
}
