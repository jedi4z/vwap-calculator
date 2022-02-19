package vwap_test

import (
	"testing"
	"vwap-calculator/vwap"

	"github.com/stretchr/testify/require"
)

func TestNewVWAPPeriod(t *testing.T) {
	t.Parallel()

	_, err := vwap.NewVWAPPeriod(200)

	require.NoError(t, err)
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		Interval     int
		Pair         string
		ExpectedVWAP float64
		Description  string
		DataPoints   []vwap.DataPoint
	}{
		{
			Description:  "should calculate the VWAP with at least 1 data point",
			Interval:     3,
			Pair:         "BTC-USD",
			ExpectedVWAP: 40178.14,
			DataPoints: []vwap.DataPoint{
				{
					Pair:   "BTC-USD",
					Price:  40178.14,
					Volume: 0.00000349,
				},
			},
		},
		{
			Description:  "should calculate the VWAP for all data points",
			Interval:     3,
			Pair:         "BTC-USD",
			ExpectedVWAP: 40177.71811124718,
			DataPoints: []vwap.DataPoint{
				{
					Pair:   "BTC-USD",
					Price:  40178.14,
					Volume: 0.00000349,
				},
				{
					Pair:   "BTC-USD",
					Price:  40177.72,
					Volume: 0.00069775,
				},
				{
					Pair:   "BTC-USD",
					Price:  40177.71,
					Volume: 0.000344,
				},
			},
		},
		{
			Description:  "should calculate the VWAP for the last 3 data points",
			Interval:     3,
			Pair:         "BTC-USD",
			ExpectedVWAP: 40177.627436234485,
			DataPoints: []vwap.DataPoint{
				{
					Pair:   "BTC-USD",
					Price:  40172.14,
					Volume: 0.00000449,
				},
				{
					Pair:   "BTC-USD",
					Price:  40178.14,
					Volume: 0.00000349,
				},
				{
					Pair:   "BTC-USD",
					Price:  40175.14,
					Volume: 0.00000329,
				},
				{
					Pair:   "BTC-USD",
					Price:  40177.72,
					Volume: 0.00069775,
				},
				{
					Pair:   "BTC-USD",
					Price:  40177.71,
					Volume: 0.000344,
				},
				{
					Pair:   "BTC-USD",
					Price:  40177.35,
					Volume: 0.00033517,
				},
			},
		},
	}

	for _, tc := range tests {
		vwapPeriod, err := vwap.NewVWAPPeriod(tc.Interval)
		require.NoError(t, err)

		var actualVWAP vwap.SumSet
		for _, d := range tc.DataPoints {
			actualVWAP = vwapPeriod.Calculate(d)
		}

		require.Equal(t, tc.ExpectedVWAP, actualVWAP[tc.Pair], tc.Description)
	}
}
