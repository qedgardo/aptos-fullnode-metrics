# Binary name
BINARY_NAME=aptos-metrics-exporter
# Compressed binary name
COMPRESSED_NAME=ame-linux-amd64
# Output directories
BUILD_DIR=build
BIN_DIR=$(BUILD_DIR)/bin

# Commands
GO_BUILD=go build
GO_GET=go get
TAR=tar
GZIP=gzip

# Install required Go dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	$(GO_GET) github.com/prometheus/client_golang/prometheus
	$(GO_GET) github.com/prometheus/client_golang/prometheus/promhttp

# Build binary for Linux amd64
.PHONY: build
build: deps
	@echo "Building the binary for Linux amd64..."
	mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(BIN_DIR)/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

# Compress the binary with gzip
.PHONY: compress-gzip
compress-gzip: build
	@echo "Compressing the binary with gzip..."
	$(TAR) -cvzf $(BUILD_DIR)/$(COMPRESSED_NAME).tar.gz -C $(BIN_DIR) $(BINARY_NAME)

# Default target (build)
.PHONY: all
all: build
