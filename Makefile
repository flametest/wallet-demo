.PHONY: gen_proto
gen_proto:
	@protoc -I ./proto --go_out=paths=source_relative:./proto --go-grpc_out=paths=source_relative:./proto ./proto/*.proto

.PHONY: build
build:
	@CGO_ENABLED=0 GOOS=linux go build -o ./cmd/wallet-demo/wallet-demo ./cmd/wallet-demo

.PHONY: run
run:
	@go run cmd/wallet-demo/main.go