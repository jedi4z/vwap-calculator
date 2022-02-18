package vwap

import (
	"sync"

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
	mu         sync.Mutex
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

func (v *vwapPeriod) GetVWAP() sumSet {
	return v.VWAP
}

func (v *vwapPeriod) Calculate(d dataPoint) {
	v.mu.Lock()
	defer v.mu.Unlock()

	// collecting datapoints by pair
	v.DataPoints[d.Pair] = append(v.DataPoints[d.Pair], d)

	// if the number of data points exceeds the interval, the first data point is removed
	if len(v.DataPoints[d.Pair]) > int(v.Interval) {
		d := v.DataPoints[d.Pair][0]
		v.DataPoints[d.Pair] = v.DataPoints[d.Pair][1:]

		v.SumPrice[d.Pair] = v.SumPrice[d.Pair] - d.Price
		v.SumVolume[d.Pair] = v.SumVolume[d.Pair] - d.Volume
	}

	if len(v.DataPoints[d.Pair]) == int(v.Interval) {
		v.SumPrice[d.Pair] = d.Price
		v.SumVolume[d.Pair] = d.Volume
		v.VWAP[d.Pair] = (d.Price * d.Volume) / d.Volume
	}
}
