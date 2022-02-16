package vwap

import "context"

type Service interface {
	Run(context.Context) error
}
