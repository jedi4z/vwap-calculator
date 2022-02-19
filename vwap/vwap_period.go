package vwap

import (
	"sync"

	"golang.org/x/xerrors"
)

type vwapPeriod struct {
	mu                sync.Mutex
	interval          int
	dataPoints        DataPointSet
	sumPrice          SumSet
	sumVolume         SumSet
	sumVolumeWeighted SumSet
	vwap              SumSet
}

func NewVWAPPeriod(interval int) (VWAPPeriod, error) {
	if interval < 0 {
		return &vwapPeriod{}, xerrors.New("the interval should be greater than 0")
	}

	return &vwapPeriod{
		interval:          interval,
		dataPoints:        make(DataPointSet),
		sumPrice:          make(SumSet),
		sumVolume:         make(SumSet),
		sumVolumeWeighted: make(SumSet),
		vwap:              make(SumSet),
	}, nil
}

func (v *vwapPeriod) Calculate(d DataPoint) SumSet {
	v.mu.Lock()
	defer v.mu.Unlock()

	// collecting datapoints by pair
	v.dataPoints[d.Pair] = append(v.dataPoints[d.Pair], d)

	// if the number of data points exceeds the interval, the first data point is removed
	if len(v.dataPoints[d.Pair]) > int(v.interval) {
		d := v.dataPoints[d.Pair][0]
		v.dataPoints[d.Pair] = v.dataPoints[d.Pair][1:]

		v.sumPrice[d.Pair] -= d.Price
		v.sumVolume[d.Pair] -= d.Volume
		v.sumVolumeWeighted[d.Pair] -= d.Price * d.Volume
	}

	v.sumPrice[d.Pair] += d.Price
	v.sumVolume[d.Pair] += d.Volume
	v.sumVolumeWeighted[d.Pair] += d.Price * d.Volume

	v.vwap[d.Pair] = v.sumVolumeWeighted[d.Pair] / v.sumVolume[d.Pair]

	return v.vwap
}
