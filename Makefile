.PHONY: run build test clean install dev

# Default target
all: build

# Install dependencies
install:
	go mod download
	go mod verify

# Run the application
run:
	go run main.go

# Build the application
build:
	go build -o bin/rideqwik-api main.go

# Build for multiple platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -o bin/rideqwik-api-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/rideqwik-api-windows-amd64.exe main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/rideqwik-api-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/rideqwik-api-darwin-arm64 main.go

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Run linter
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out

# Run with hot reload (requires air)
dev:
	air

# Generate swagger documentation (requires swag)
swagger:
	swag init

# Install development tools
install-tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

