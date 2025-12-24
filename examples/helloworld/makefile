# 定义目标路径和源路径
GOROOT := $(shell go env GOROOT)
WASM_EXEC_JS_SRC = $(GOROOT)/lib/wasm/wasm_exec.js
DEST_JS = static/js/wasm_exec.js
CURRENT_DIR := $(notdir $(CURDIR))

# 默认目标
.PHONY: all
all: dev

init:
	@mkdir -p static
	@mkdir -p static/js
	@cp $(WASM_EXEC_JS_SRC) $(DEST_JS)

build:
	@GOARCH=wasm GOOS=js go build -o ./static/app.wasm main.go

# 运行程序
run: build
	@go build
	@echo "Build complete."

dev: build
	@go build -o ./ cmd/main.go
	@echo "Build complete."
	@./main

clean:
	@go clean ./...
	@-rm static/app.wasm