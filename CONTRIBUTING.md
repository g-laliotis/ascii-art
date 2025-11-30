# Contributing to ASCII-Art

Thank you for your interest in contributing to ASCII-Art! This document provides guidelines for contributing to the project.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/g-laliotis/ascii-art.git`
3. Create a feature branch: `git checkout -b feature-name`

## Development Setup

```bash
# Install Go 1.19 or higher
# Clone the repository
cd ascii-art
go mod tidy
```

## Making Changes

1. **Code Style**: Follow Go conventions and run `go fmt`
2. **Testing**: Add tests for new features and ensure all tests pass
3. **Documentation**: Update README.md if adding new features

## Testing

```bash
# Run all tests
make test

# Run with coverage
make test-coverage

# Format code
make fmt
```

## Submitting Changes

1. Ensure all tests pass
2. Commit with clear, descriptive messages
3. Push to your fork
4. Create a Pull Request

## Adding New Banner Styles

To add new banner styles (like `shadow` or `thinkertoy`):

1. Add banner file to `assets/` directory
2. Update banner loading logic in `internal/ascii/banner.go`
3. Add comprehensive tests
4. Update documentation

## Code Review Process

- All submissions require review
- Tests must pass
- Code must follow Go best practices
- Documentation must be updated

## Questions?

Open an issue for questions or discussions about contributions.