# Makefile for the BillGame project

# Define source and output variables
CLIENT_SRC := ./cmd/client
SERVER_SRC := ./cmd/server
CLIENT_BIN := ./bin/client
SERVER_BIN := ./bin/server

# The default command, executed when you just run "make"
.PHONY: all
all: build

# Builds all application binaries
.PHONY: build
build: build-client build-server

# Builds the client binary
.PHONY: build-client
build-client:
	@echo "Building client..."
	@go build -o $(CLIENT_BIN) $(CLIENT_SRC)

# Builds the server binary
.PHONY: build-server
build-server:
	@echo "Building server..."
	@go build -o $(SERVER_BIN) $(SERVER_SRC)

# Removes the build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf ./bin/*
