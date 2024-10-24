package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/qedgardo/aptos-metrics-exporter/internal/collector"
)

func main() {
	// -p flag for the server port. Default value: 2112
	var port int

	flag.IntVar(&port, "p", 2112, "Server port")
	flag.Parse()

	// Set up the HTTP handler for Prometheus metrics
	http.Handle("/metrics", promhttp.Handler())

	// Create a ticker that will call FetchLatestBlockHeight every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for range ticker.C {
			collector.FetchLatestBlockHeight()
		}
	}()

	// Fetch the block height initially to set the initial value before scraping starts
	collector.FetchLatestBlockHeight()

	address := fmt.Sprintf(":%d", port)
	// Start the HTTP server
	log.Printf("Starting server on %s", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
