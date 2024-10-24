package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/qedgardo/aptos-metrics-exporter/internal/collector"
)

func main() {
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

	// Start the HTTP server
	log.Println("Starting server on :2112")
	if err := http.ListenAndServe(":2112", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
