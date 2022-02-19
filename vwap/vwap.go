package vwap

import "context"

type DataPointSet map[string][]DataPoint

type SumSet map[string]float64

type DataPoint struct {
	Pair   string
	Price  float64
	Volume float64
}

type Service interface {
	Run(context.Context) error
}

type VWAPPeriod interface {
	Calculate(DataPoint)
	GetVWAP() SumSet
}
