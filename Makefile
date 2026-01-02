# Binary name
BINARY := rpt

# Go settings
GO := go
GOFLAGS :=
LDFLAGS := -ldflags "-X main.Version=dev"

# Tools
GOLANGCI_LINT := $(shell go env GOPATH)/bin/golangci-lint

.PHONY: all build test clean install fmt lint lint-install

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

## Install golangci-lint if not present
lint-install:
	@if [ ! -x "$(GOLANGCI_LINT)" ]; then \
		echo "==> Installing golangci-lint"; \
		$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2; \
	fi

## Run linter
lint: lint-install
	@echo "==> Running golangci-lint"
	$(GOLANGCI_LINT) run ./...

## Clean build artifacts
clean:
	@echo "==> Cleaning"
	rm -f $(BINARY)
	rm -rf build
