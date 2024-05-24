build:
	@go build -o bin/api-ecommerce cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/api-ecommerce