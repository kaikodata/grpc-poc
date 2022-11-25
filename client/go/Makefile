.PHONY: generate-bindings
generate-bindings:
	protoc -I=../.. ../../kaiko.proto --go_out=plugins=grpc:./kaiko

.PHONY: build-client
build-client:
	go build -o kaiko-exists main.go
