APP_NAME=obsidianOptimizeMCP
SRC_DIR=.
BIN_DIR=bin
MAIN=main.go

.PHONY: build run clean test

build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(MAIN)

run: build
	@echo "Running $(APP_NAME)..."
	@$(BIN_DIR)/$(APP_NAME)

clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR)

fmt:
	gofmt -w $(SRC_DIR)

vet:
	go vet ./...

test:
	go test ./... -v
