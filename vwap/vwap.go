package vwap

import "context"

type Service interface {
	Run(context.Context) error
}

type VWAPPeriod interface {
	Calculate(dataPoint)
	GetVWAP() sumSet
}
