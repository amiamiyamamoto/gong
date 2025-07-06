.PHONY: build run clean test help

# Default target
all: build

# Build the application
build:
	go build -o gong main.go

# Run the application
run:
	go run main.go

# Clean build artifacts
clean:
	rm -f gong

# Run tests
test:
	go test ./...

# Run tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Install dependencies
deps:
	go mod tidy
	go mod download

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Show help
help:
	@echo "Available targets:"
	@echo "  build         Build the application"
	@echo "  run           Run the application"
	@echo "  clean         Clean build artifacts"
	@echo "  test          Run tests"
	@echo "  test-coverage Run tests with coverage report"
	@echo "  deps          Install/update dependencies"
	@echo "  fmt           Format code"
	@echo "  lint          Lint code"
	@echo "  help          Show this help message"
