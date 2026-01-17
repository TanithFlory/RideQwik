.PHONY: run build test clean install dev sqlc

all: build

install:
	go mod download
	go mod verify

sqlc:
	sqlc generate

run:
	go run main.go

build:
	go build -o bin/rideqwik-api main.go

build-all:
	GOOS=linux GOARCH=amd64 go build -o bin/rideqwik-api-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/rideqwik-api-windows-amd64.exe main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/rideqwik-api-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/rideqwik-api-darwin-arm64 main.go

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

lint:
	golangci-lint run

fmt:
	go fmt ./...

clean:
	rm -rf bin/
	rm -f coverage.out

dev:
	air

swagger:
	swag init

install-tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
