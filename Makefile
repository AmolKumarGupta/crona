BINARY_NAME=crona
SRC=main.go
BIN_DIR=bin
GOPATH=$(shell go env GOPATH)

.PHONY: all build run clean

# Default target, builds the Go binary
all: build


build:
	@echo "Building the binary..."
	@go build -o $(BIN_DIR)/$(BINARY_NAME) $(SRC)

run:
	@$(BIN_DIR)/$(BINARY_NAME)

clean:
	@echo "Cleaning the bin directory..."
	@rm $(BIN_DIR)/$(BINARY_NAME)

lint:
	@$(GOPATH)/bin/golangci-lint run