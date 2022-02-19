package vwap

import (
	"sync"

	"golang.org/x/xerrors"
)

type dataPointSet map[string][]dataPoint

type sumSet map[string]float64

type dataPoint struct {
	pair   string
	price  float64
	volume float64
}

type vwapPeriod struct {
	mu         sync.Mutex
	interval   int
	dataPoints dataPointSet
	sumPrice   sumSet
	sumVolume  sumSet
	vwap       sumSet
}

func NewVWAPPeriod(interval int) (VWAPPeriod, error) {
	if interval < 0 {
		return &vwapPeriod{}, xerrors.New("the interval should be greater than 0")
	}

	return &vwapPeriod{
		interval:   interval,
		dataPoints: make(dataPointSet),
		sumPrice:   make(sumSet),
		sumVolume:  make(sumSet),
		vwap:       make(sumSet),
	}, nil
}

func (v *vwapPeriod) GetVWAP() sumSet {
	return v.vwap
}

func (v *vwapPeriod) Calculate(d dataPoint) {
	v.mu.Lock()
	defer v.mu.Unlock()

	// collecting datapoints by pair
	v.dataPoints[d.pair] = append(v.dataPoints[d.pair], d)

	// if the number of data points exceeds the interval, the first data point is removed
	if len(v.dataPoints[d.pair]) > int(v.interval) {
		d := v.dataPoints[d.pair][0]
		v.dataPoints[d.pair] = v.dataPoints[d.pair][1:]

		v.sumPrice[d.pair] = v.sumPrice[d.pair] - d.price
		v.sumVolume[d.pair] = v.sumVolume[d.pair] - d.volume
	}

	if len(v.dataPoints[d.pair]) == int(v.interval) {
		v.sumPrice[d.pair] = d.price
		v.sumVolume[d.pair] = d.volume
		v.vwap[d.pair] = (d.price * d.volume) / d.volume
	}
}
