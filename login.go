package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type LoggingService struct {
	next PriceFetcher // next is used to call the next service in the chain
}

func NewLoggingService(next PriceFetcher) PriceFetcher {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	// s is the logging service used as a wrapper for the next service in the chain in function it is used as a wrapper , wrapper is used to add additional functionality to the function
	// logrus is used to log the time taken by the function to execute
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{ // logrus.Fields is used to add additional fields to the log
			"requestID": ctx.Value("requestID"), // ctx.Value("requestID") is used to get the value of requestID from the context
			"took":      time.Since(begin),      // time.Since(begin) is used to calculate the time taken by the function to execute
			"err":       err,
			"price":     price,
		}).Info("fetchPrice") // .Info is used to log the info,fetchPrice is the name of the function , .Info("fetchPrice") is used to log the info of the function
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker) // s.next.FetchPrice(ctx,ticker) is used to call the next service in the chain
}
