# Development Guide

## Getting Started

### Environment Setup

1. **Go Installation**
   ```bash
   # Verify Go installation
   go version  # Should be 1.19+
   ```

2. **Project Initialization**
   ```bash
   cd ascii-art
   go mod init ascii-art
   go mod tidy
   ```

### Development Workflow

#### 1. Create Directory Structure
```bash
mkdir -p cmd/ascii-art internal/ascii internal/version assets testdata/samples testdata/expected
```

#### 2. Development Order
1. **Setup Phase**: Create `go.mod`, `cmd/ascii-art/main.go`, `internal/version/version.go`
2. **Banner Phase**: Implement banner loading (`internal/ascii/banner.go`) with tests
3. **Core Phase**: Implement ASCII generation (`internal/ascii/art.go`) with tests
4. **Testing Phase**: Write comprehensive unit tests in `*_test.go` files
5. **Integration Phase**: Connect all components and finalize

### Implementation Guidelines

#### Code Style
- Follow `gofmt` formatting
- Use meaningful variable names
- Keep functions under 50 lines
- Add comments for complex logic
- Use Go naming conventions

#### Error Handling Pattern
```go
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}
```

#### Testing Pattern
```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
        wantErr  bool
    }{
        // test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

### Key Implementation Details

#### Banner File Processing
- Each character occupies exactly 8 lines
- Characters 32-126 (95 total characters)
- Parse sequentially: space, !, ", #, ..., ~

#### String Processing
- Handle `\n` as line breaks
- Process each line independently
- Combine character representations horizontally

#### Output Format
- Each output line corresponds to one of 8 character lines
- Maintain exact spacing and alignment
- End with newline character

### Development Commands

#### Building
```bash
# Development build
go build .

# Optimized build
go build -ldflags="-s -w" .
```

#### Testing
```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test -run TestSpecificFunction ./tests
```

#### Code Quality
```bash
# Format code
go fmt ./...

# Vet code
go vet ./...

# Run linter (if installed)
golangci-lint run
```

### Debugging Tips

#### Common Issues
1. **Banner Loading**: Verify file paths and format
2. **Character Mapping**: Check ASCII value calculations
3. **Output Formatting**: Validate line endings and spacing

#### Debug Techniques
- Use `fmt.Printf` for intermediate values
- Test with simple inputs first
- Validate banner file parsing separately

### Performance Optimization

#### Memory Efficiency
- Pre-allocate slices with known capacity
- Use `strings.Builder` for output construction
- Cache banner data after loading

#### Time Efficiency
- Minimize string operations
- Use efficient data structures
- Avoid unnecessary allocations

### Git Workflow

#### Commit Strategy
```bash
# Feature development
git checkout -b feature/banner-loading
git add .
git commit -m "feat: implement banner file loading"

# Testing
git checkout -b test/banner-tests
git add .
git commit -m "test: add banner loading tests"
```

#### Commit Message Format
- `feat:` for new features
- `fix:` for bug fixes
- `test:` for adding tests
- `docs:` for documentation
- `refactor:` for code refactoring

### Quality Checklist

Before committing:
- [ ] Code compiles without errors
- [ ] All tests pass
- [ ] Code is formatted (`go fmt`)
- [ ] No linter warnings
- [ ] Documentation updated
- [ ] Edge cases handled

### Deployment Preparation

#### Final Build
```bash
# Create optimized binary
go build -ldflags="-s -w" -o ascii-art .

# Test with examples
./ascii-art "Hello"
./ascii-art "Hello\nWorld"
```

#### Submission Checklist
- [ ] All required functionality implemented
- [ ] Comprehensive tests written
- [ ] Code follows Go best practices
- [ ] Documentation complete
- [ ] No external dependencies
- [ ] Handles all specified edge cases

This guide provides a structured approach to developing your ASCII-art project efficiently and maintainably.