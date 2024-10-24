package collector

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	aptosLatestBlockHeight = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "aptos_latest_block_height",
		Help: "The latest block height of the Aptos blockchain.",
	})
)

func init() {
	prometheus.MustRegister(aptosLatestBlockHeight)
}

// FetchLatestBlockHeight retrieves the latest block height from the Aptos node.
func FetchLatestBlockHeight() {
	// Replace with your Aptos node's endpoint
	resp, err := http.Get("http://localhost:8080/v1")
	if err != nil {
		log.Println("Error fetching latest block height:", err)
		return
	}
	defer resp.Body.Close()

	var result struct {
		BlockHeight string `json:"block_height"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}

	// Convert the block height string to int64 to avoid scientific notation
	blockHeight, err := strconv.ParseInt(result.BlockHeight, 10, 64)
	if err != nil {
		log.Println("Error converting block height to int64:", err)
		return
	}

	// Set the block height in the gauge as a float64, but still as an integer value
	aptosLatestBlockHeight.Set(float64(blockHeight))
}
