package vwap

import (
	"golang.org/x/xerrors"
)

type dataPoint struct {
	Pair   string
	Price  float64
	Volume float64
}

type vwapPeriod struct {
	Interval   uint
	DataPoints []dataPoint
	SumPrice   map[string]float64
	SumVolume  map[string]float64
	VWAP       map[string]float64
}

func NewVWAPPeriod(interval uint) (VWAPPeriod, error) {
	if interval < 0 {
		return &vwapPeriod{}, xerrors.New("the interval should be greater than 0")
	}

	return &vwapPeriod{
		Interval:   interval,
		DataPoints: []dataPoint{},
		SumPrice:   make(map[string]float64),
		SumVolume:  make(map[string]float64),
		VWAP:       make(map[string]float64),
	}, nil
}

func (v *vwapPeriod) Calculate(d dataPoint) {
	v.VWAP["test"] += 1
}

func (v vwapPeriod) GetVWAP() map[string]float64 {
	return v.VWAP
}
