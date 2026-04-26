# Project variables
BINARY_NAME=e-commerce
MAIN_PACKAGE_PATH=./cmd/main.go
BUILD_DIR=bin

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOGEN=$(GOCMD) generate

.PHONY: all build run test clean generate tidy help

# Default target
all: generate build

## generate: Run go generate to trigger //go:generate directives
generate:
	@echo "==> Generating code..."
	$(GOGEN) ./...

## build: Compile the binary
build: generate
	@echo "==> Building binary..."
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE_PATH)

## run: Build and run the application
run: build
	@echo "==> Running application..."
	./$(BUILD_DIR)/$(BINARY_NAME)

## test: Run unit tests
test:
	@echo "==> Running tests..."
	$(GOTEST) -v -race ./...

## tidy: Clean up go.mod and go.sum
tidy:
	@echo "==> Tidying modules..."
	$(GOMOD) tidy

## clean: Remove build artifacts
clean:
	@echo "==> Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' |  sed -e 's/^/ /'
