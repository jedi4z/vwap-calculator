package vwap

import (
	"fmt"

	"golang.org/x/xerrors"
)

type dataPointSet map[string][]dataPoint

type sumSet map[string]float64

type dataPoint struct {
	Pair   string
	Price  float64
	Volume float64
}

type vwapPeriod struct {
	Interval   uint
	DataPoints dataPointSet
	SumPrice   sumSet
	SumVolume  sumSet
	VWAP       sumSet
}

func NewVWAPPeriod(interval uint) (VWAPPeriod, error) {
	if interval < 0 {
		return &vwapPeriod{}, xerrors.New("the interval should be greater than 0")
	}

	return &vwapPeriod{
		Interval:   interval,
		DataPoints: make(dataPointSet),
		SumPrice:   make(sumSet),
		SumVolume:  make(sumSet),
		VWAP:       make(sumSet),
	}, nil
}

func (v *vwapPeriod) Calculate(d dataPoint) {
	// collecting datapoints by pair
	v.DataPoints[d.Pair] = append(v.DataPoints[d.Pair], d)

	if len(v.DataPoints[d.Pair]) == int(v.Interval) {
		fmt.Printf("Data point set for pair %s, current size: %d \n", d.Pair, len(v.DataPoints[d.Pair]))
	}
}

func (v vwapPeriod) GetVWAP() map[string]float64 {
	return v.VWAP
}
