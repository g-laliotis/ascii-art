# ASCII-Art Makefile

.PHONY: build test clean install run run-shadow run-thinkertoy run-color run-output help

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
	rm -f ascii-art *.txt

# Install to GOPATH/bin
install:
	go install ./cmd/ascii-art

# Run examples
run:
	go run ./cmd/ascii-art "Hello World"

run-shadow:
	go run ./cmd/ascii-art "Hello" shadow

run-thinkertoy:
	go run ./cmd/ascii-art "Hello" thinkertoy

run-color:
	go run ./cmd/ascii-art --color=red "Hello World"

run-output:
	go run ./cmd/ascii-art --output=example.txt "Hello World"
	@echo "Output saved to example.txt"

run-all-features:
	go run ./cmd/ascii-art --color=blue --output=demo.txt "Demo" thinkertoy
	@echo "Colored thinkertoy output saved to demo.txt"

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Show help
help:
	@echo "Available targets:"
	@echo "  build            - Build the binary"
	@echo "  test             - Run all tests"
	@echo "  test-coverage    - Run tests with coverage"
	@echo "  clean            - Remove build artifacts and output files"
	@echo "  install          - Install to GOPATH/bin"
	@echo "  run              - Run with standard banner"
	@echo "  run-shadow       - Run with shadow banner"
	@echo "  run-thinkertoy   - Run with thinkertoy banner"
	@echo "  run-color        - Run with color example"
	@echo "  run-output       - Run with file output example"
	@echo "  run-all-features - Run with all features combined"
	@echo "  fmt              - Format code"
	@echo "  lint             - Lint code (requires golangci-lint)"
	@echo "  help             - Show this help"