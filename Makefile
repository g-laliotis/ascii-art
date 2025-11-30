# ASCII-Art Makefile

.PHONY: build test clean install run help

# Default target
all: build

# Build the binary
build:
	go build -o ascii-art ./cmd/ascii-art

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -cover ./...

# Clean build artifacts
clean:
	rm -f ascii-art

# Install to GOPATH/bin
install:
	go install ./cmd/ascii-art

# Run with example
run:
	go run ./cmd/ascii-art "Hello World"

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Show help
help:
	@echo "Available targets:"
	@echo "  build         - Build the binary"
	@echo "  test          - Run all tests"
	@echo "  test-coverage - Run tests with coverage"
	@echo "  clean         - Remove build artifacts"
	@echo "  install       - Install to GOPATH/bin"
	@echo "  run           - Run with example text"
	@echo "  fmt           - Format code"
	@echo "  lint          - Lint code (requires golangci-lint)"
	@echo "  help          - Show this help"