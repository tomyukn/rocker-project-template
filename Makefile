# Binary name
BINARY := rpt

# Go settings
GO := go
GOFLAGS :=
LDFLAGS := -ldflags "-X main.Version=dev"

# Directories
BUILD_DIR := build

.PHONY: all build test clean install fmt lint

all: test build

## Build the binary for local use
build:
	@echo "==> Building $(BINARY)"
	$(GO) build $(GOFLAGS) $(LDFLAGS) -o $(BINARY)

## Run tests
test:
	@echo "==> Running tests"
	$(GO) test ./...

## Install binary to GOPATH/bin
install:
	@echo "==> Installing $(BINARY)"
	$(GO) install $(LDFLAGS)

## Format Go code
fmt:
	@echo "==> Formatting code"
	$(GO) fmt ./...

## Clean build artifacts
clean:
	@echo "==> Cleaning"
	rm -f $(BINARY)
	rm -rf $(BUILD_DIR)
