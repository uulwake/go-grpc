all: 
	@echo "No default. Specify command please."

generate-proto:
	@protoc --go_out=. --go-grpc_out=. \
		./proto/hello_world/hello_world.proto \
		./proto/user/user.proto \
		./shared-proto/item/item.proto \
		./shared-proto/order/order.proto 

run-grpc-server: generate-proto
	@go run cmd/server/main.go

run-grpc-client: generate-proto
	@go run cmd/client/main.go

watch-grpc-server: generate-proto
	@nodemon --watch './**/*.go' --ignore ./generated --signal SIGTERM --exec 'go' run cmd/server/main.go

submodule-status:
	@git submodule foreach git status

submodule-pull:
	@git submodule foreach git pull origin main