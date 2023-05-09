all: 
	@echo "No default. Specify command please."

generate-proto:
	@protoc --go_out=. --go-grpc_out=. \
		./proto/hello_world/hello_world.proto

run-grpc-server: generate-proto
	@go run cmd/server/main.go

run-grpc-client: generate-proto
	@go run cmd/client/main.go