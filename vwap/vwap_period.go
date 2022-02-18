package vwap

import (
	"golang.org/x/xerrors"
)

const defaultInterval = 200

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

func NewVWAPPeriod(dataPoint []dataPoint, interval uint) (VWAPPeriod, error) {
	if interval == 0 {
		interval = defaultInterval
	}

	if len(dataPoint) > int(interval) {
		return vwapPeriod{}, xerrors.New("initial datapoints exceeds maxSize")
	}

	return vwapPeriod{
		Interval:   interval,
		DataPoints: dataPoint,
		SumPrice:   make(map[string]float64),
		SumVolume:  make(map[string]float64),
		VWAP:       make(map[string]float64),
	}, nil
}

func (v vwapPeriod) Calculate(d dataPoint) {

}
