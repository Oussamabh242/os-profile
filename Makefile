BINARY_NAME=main
BIN_DIR=bin
WASM_OUTPUT=web/main.wasm
WASM_EXEC=web/wasm_exec.js
WEB_SERVER_BIN=web
WEB_SERVER_DIR=web

.PHONY: all
all: build

.PHONY: build
build:
	go build -o $(BIN_DIR)/$(BINARY_NAME) ./cmd/engine.go

.PHONY: build-wasm
build-wasm:
	GOOS=js GOARCH=wasm go build -o $(WASM_OUTPUT) ./cmd/engine.go

.PHONY: run
run: build
	./$(BIN_DIR)/$(BINARY_NAME)

.PHONY: run-wasm
run-wasm:
	go run ./bin/web

.PHONY: fmt
fmt:
	go fmt ./...
