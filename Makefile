.PHONY: proto build run

proto:
	protoc --go_out=./src/proto-gen --go_opt=paths=source_relative --go-grpc_out=./src/proto-gen --go-grpc_opt=paths=source_relative ./proto/device-service.proto

build:
	go build -o device-provider ./cli/device-provider/main.go

run-server:
	go run ./cli/device_provider/main.go -port 50051 -logtostderr

list-device:
	go run ./cli/device_consumer/main.go -address localhost:50051 list
