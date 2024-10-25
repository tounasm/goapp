.DEFAULT_GOAL := all

BIN_DIR := bin
SERVER := ./cmd/server/main.go
CLIENT := ./cmd/client/main.go

.PHONY: all
all: clean server client test

.PHONY: server
server:
	@mkdir -p $(BIN_DIR)
	@echo "Building server..."
	go build -o $(BIN_DIR)/server $(SERVER)

.PHONY: client
client:
	@mkdir -p $(BIN_DIR)
	@echo "Building client..."
	go build -o $(BIN_DIR)/client $(CLIENT)

.PHONY: clean
clean:
	@echo "Cleaning up..."
	go clean
	rm -f $(BIN_DIR)/*

.PHONY: test
test:
	@echo "Running tests..."
	go test -v ./... -cover

.PHONY: bench
bench:
	@echo "Running benchmarks..."
	go test -v ./... -bench=. -run=xxx -benchmem