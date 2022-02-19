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
		VWAPExpected float64
		DataPoints   []vwap.DataPoint
	}{
		{
			Interval:     3,
			Pair:         "BTC-USD",
			VWAPExpected: 40177.71,
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
		// {
		// 	Interval:     3,
		// 	Pair:         "BTC-USD",
		// 	VWAPExpected: 40177.62,
		// 	DataPoints: []vwap.DataPoint{
		// 		{
		// 			Pair:   "BTC-USD",
		// 			Price:  40178.14,
		// 			Volume: 0.00000349,
		// 		},
		// 		{
		// 			Pair:   "BTC-USD",
		// 			Price:  40177.72,
		// 			Volume: 0.00069775,
		// 		},
		// 		{
		// 			Pair:   "BTC-USD",
		// 			Price:  40177.71,
		// 			Volume: 0.000344,
		// 		},
		// 		{
		// 			Pair:   "BTC-USD",
		// 			Price:  40177.35,
		// 			Volume: 0.00033517,
		// 		},
		// 	},
		// },
	}

	for _, tc := range tests {
		vwapPeriod, err := vwap.NewVWAPPeriod(tc.Interval)
		require.NoError(t, err)

		for _, d := range tc.DataPoints {
			vwapPeriod.Calculate(d)

		}

		require.Equal(t, tc.VWAPExpected, vwapPeriod.GetVWAP()[tc.Pair])
	}
}
