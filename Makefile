build:
	@go build -o bin/Product-Management

run:build
	@./bin/Product-Management

test:
	@go test -v ./...
