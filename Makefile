build:
	@go build -o ./bin/shop
run: build
	@go run shop
generate:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/*.proto

.PHONY: proto