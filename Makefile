.PHONY: gen_proto
	@protoc -I ./proto --go_out=paths=source_relative:./proto --go-grpc_out=paths=source_relative:./proto ./proto/*.proto