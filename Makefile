build:
	@go build -o bin/api-ecommerce cmd/main.go

test:
	@go test -V ./,,,

run: build
	@./bin/api-ecommerce