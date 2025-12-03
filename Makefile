.PHONY: build clean test lint install release help

# Build variables
BINARY_NAME=kube-dependency-checker
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
GIT_COMMIT?=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE?=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS=-ldflags "-X github.com/pmady/kube-dependency-checker/cmd.Version=$(VERSION) \
	-X github.com/pmady/kube-dependency-checker/cmd.GitCommit=$(GIT_COMMIT) \
	-X github.com/pmady/kube-dependency-checker/cmd.BuildDate=$(BUILD_DATE)"

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Default target
all: build

## build: Build the binary
build:
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) .

## build-all: Build for all platforms
build-all: build-linux build-darwin build-windows

build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-arm64 .

build-darwin:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64 .

build-windows:
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe .

## test: Run tests
test:
	$(GOTEST) -v ./...

## test-coverage: Run tests with coverage
test-coverage:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

## lint: Run linter
lint:
	golangci-lint run

## clean: Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf dist/

## install: Install the binary
install: build
	mv $(BINARY_NAME) $(GOPATH)/bin/

## tidy: Tidy go modules
tidy:
	$(GOMOD) tidy

## release: Create a new release (requires goreleaser)
release:
	goreleaser release --clean

## release-snapshot: Create a snapshot release (for testing)
release-snapshot:
	goreleaser release --snapshot --clean

## help: Show this help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'
