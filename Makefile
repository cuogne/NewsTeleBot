.PHONY: build run dev test clean fmt lint deps help

help:
	@echo "Available commands:"
	@echo "  make setup       - Setup development environment"
	@echo "  make build       - Build the bot binary"
	@echo "  make run         - Run the bot directly"
	@echo "  make dev         - Run the bot with hot reload (requires air)"
	@echo "  make test        - Run all tests"
	@echo "  make fmt         - Format code"
	@echo "  make lint        - Run linter (requires golangci-lint)"
	@echo "  make deps        - Download and tidy dependencies"
	@echo "  make clean       - Remove build artifacts"

setup:
	@echo "setup make air for hot reloading..."
	@grep -q '$$PATH:$$(go env GOPATH)/bin' ~/.zshrc || echo 'export PATH=$$PATH:$$(go env GOPATH)/bin' >> ~/.zshrc
	@source ~/.zshrc
	@go install github.com/air-verse/air@latest
	@echo "Setup complete!"

build:
	go build -o ./bin/bot ./cmd/bot

run:
	go run ./cmd/bot

dev:
	air

test:
	go test -v ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

deps:
	go mod download
	go mod tidy

clean:
	rm -rf ./bin
