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
4. **Cross-platform**: Consider Unix/Windows compatibility for system calls

## Testing

```bash
# Run all tests
make test

# Run with coverage
make test-coverage

# Test different banners
make run-shadow
make run-thinkertoy

# Test file output
make run-output

# Test all features
make run-all-features

# Format code
make fmt
```

## Adding New Features

### Adding New Banner Styles

To add new banner styles:

1. Create banner file in `assets/` directory (855 lines: 8 lines per character + separators)
2. Each character must be exactly 8 lines tall
3. Include all ASCII printable characters (32-126)
4. Add tests in `internal/ascii/art_banner_test.go`
5. Update documentation

### Adding New Color Support

1. Add color codes to `colorMap` in `internal/ascii/color.go`
2. Add tests in `internal/ascii/color_test.go`
3. Update usage documentation

### Adding Cross-platform Features

1. Use build tags for OS-specific code (`//go:build unix` or `//go:build windows`)
2. Create separate files for different platforms (e.g., `terminal_unix.go`, `terminal_windows.go`)
3. Test on multiple platforms

## File Structure Guidelines

- `cmd/`: Entry points and CLI handling
- `internal/ascii/`: Core ASCII art logic
- `assets/`: Banner template files
- Tests should be comprehensive and cover edge cases
- OS-specific code should use build tags

## Submitting Changes

1. Ensure all tests pass (`make test`)
2. Run code formatting (`make fmt`)
3. Test on multiple platforms if applicable
4. Commit with clear, descriptive messages
5. Push to your fork
6. Create a Pull Request with detailed description

## Code Review Process

- All submissions require review
- Tests must pass on all platforms
- Code must follow Go best practices
- Documentation must be updated
- New features need comprehensive tests

## Performance Considerations

- Keep memory usage minimal
- Avoid external dependencies
- Optimize for terminal width detection
- Consider large text input handling

## Questions?

Open an issue for questions or discussions about contributions.