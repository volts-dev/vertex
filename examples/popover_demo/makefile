# Definitions
GOROOT := $(shell go env GOROOT)
WASM_EXEC_JS_SRC = $(GOROOT)/lib/wasm/wasm_exec.js
DEST_JS = static/js/wasm_exec.js

.PHONY: all init build run clean

all: run

init:
	@mkdir -p static
	@mkdir -p static/js
	@cp $(WASM_EXEC_JS_SRC) $(DEST_JS)

build:
	@GOARCH=wasm GOOS=js go build -o ./static/app.wasm main.go

run: build
	@go run cmd/main.go

clean:
	@rm -rf static
