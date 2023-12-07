package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next PriceFetcher
}

func NewMetricService(next PriceFetcher) PriceFetcher {
	return &metricService{next: next}
}

func (s *metricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("pushing metrics to prometheus") // prometheus is a metrics server that can be used to monitor the health of your application
	return s.next.FetchPrice(ctx, ticker)
}
