.PHONY: build-server
build-server:
	go build -o kaiko-server main.go

.PHONY: build-grpc
generate-bindings:
	protoc --go-grpc_out=../ --go-grpc_opt=paths=source_relative --proto_path=../ kaiko.proto
