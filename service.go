package main

import (
	"context"
	"fmt"
	"time"
)

// interface that can fetch a price for a given ticker // TICKER IS for eg: BTC, ETH
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error) // take context  and string for price, return float64 and error
}


// implements the interface wala pricefetcher
type priceFetcher struct {
}

// this func is buisness logic which is required in microservices // buisness logic should be clean and no json references idhar
func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx,ticker)
}

var priceMocks = map[string]float64{
	"BTC": 10_000.0,
	"ETH": 500.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	time.Sleep(100*time.Millisecond)
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}

	return price ,nil
}
