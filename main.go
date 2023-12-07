package main

import (
	// "context"
	"flag"
	// "fmt"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3100", "listen address the service is working")
	flag.Parse()
	svc := NewLoggingService(NewMetricService(&priceFetcher{}))
	// price, err := svc.FetchPrice(context.Background(), "ETH")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
	// fmt.Println(price)
}
