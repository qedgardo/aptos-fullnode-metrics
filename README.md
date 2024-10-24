# Aptos Metrics Exporter

Atm **Aptos Metrics Exporter** is a Prometheus exporter that collects and exposes the latest block height in a prometheus format metric from an Aptos blockchain node. 
By default, the exporter will fetch the block height in `http://localhost:8080/v1` and start an HTTP server on port `2112` to expose the metric.

## Prerequisites

- **Go**: Make sure you have Go installed to build the project.
- **Prometheus**: The exporter is designed to be scraped by a Prometheus instance.

## Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/qedgardo/aptos-metrics-exporter.git
    cd aptos-metrics-exporter
    ```

2. **Build the binary**:

    ```bash
    make build
    ```
    You can find the built binary in the `build/bin` directory.
3. **Run the exporter**:

    After building the binary, run it:

    ```bash
    ./aptos-metrics-exporter
    ```

    You can specify a different port by adding the `-p` flag followed by the desired port number. For example, to run the exporter on port `9090`:
    ```bash
    ./aptos-metrics-exporter -p 9090
    ```

4. **Prometheus Configuration**:

    Add the following scrape job to your Prometheus configuration (`prometheus.yml`):

    ```yaml
    scrape_configs:
      - job_name: 'aptos-metrics-exporter'
        static_configs:
          - targets: ['localhost:2112']
    ```

    Make sure to replace `localhost` with the actual IP address or domain where the exporter is running if it's on a remote machine.

## Usage

Once the exporter is running, it will expose metrics at the `/metrics` endpoint. Prometheus can then scrape this endpoint at regular intervals to collect metrics.

### Exposed Metrics

- **`aptos_latest_block_height`**: The current/latest block height of the Aptos blockchain node.

Example output from `/metrics`:

```
# HELP aptos_latest_block_height The latest block height of the Aptos blockchain.
# TYPE aptos_latest_block_height gauge
aptos_latest_block_height 28832207
```

### Testing

To test the project locally, you can use `curl` to query the `/metrics` endpoint:

```bash
curl http://localhost:2112/metrics
```

