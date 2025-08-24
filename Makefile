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
	gzip -k -9 web/main.wasm

.PHONY: run
run: build
	./$(BIN_DIR)/$(BINARY_NAME)

.PHONY: run-wasm
run-wasm: build-wasm
	./bin/web

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build-web
build-web:
	go build -o ./bin/web ./cmd/web.go

web: build-web run-wasm



test:
	@all_pkgs=$$(mktemp); ignored_pkgs=$$(mktemp); filtered_pkgs=$$(mktemp); \
	go list ./... | sort > $$all_pkgs; \
	cat .testignore | xargs -n1 go list 2>/dev/null | sort > $$ignored_pkgs; \
	comm -23 $$all_pkgs $$ignored_pkgs > $$filtered_pkgs; \
	go test -v $$(cat $$filtered_pkgs); \
	rm -f $$all_pkgs $$ignored_pkgs $$filtered_pkgs

