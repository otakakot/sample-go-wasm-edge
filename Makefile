.PHONY: build
build: ## go build
	@go mod tidy
	@go mod vendor
	@GOOS=wasip1 GOARCH=wasm go build -o main.wasm main.go

.PHONY: run
run: ## execute main.wasm
	@wasmedge main.wasm
