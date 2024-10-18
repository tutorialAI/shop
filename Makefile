build:
	@go build -o ./bin/shop
run: build
	@go run shop
