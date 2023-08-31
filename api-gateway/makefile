SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin

build: ${BINARY_DIR} ## Compile the code, build Executable File
	$(GOCMD) build -o api -v ./cmd/api

run: ## Start application
	$(GOCMD) run ./cmd

test: ## Run tests
	$(GOCMD) test ./... -cover

wire: ## Generate wire_gen.go
	cd pkg/di && wire

swag: ## Generate swagger docs
	swag init -g pkg/api/handler/admin.go -o ./cmd/api/docs

proto:
	protoc pkg/proto/**/*.proto --go_out=. --go-grpc_out=.
