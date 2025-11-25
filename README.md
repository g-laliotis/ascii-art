# ASCII-Art

A Go program that converts text strings into ASCII art using predefined banner templates.

## Overview

ASCII-Art receives a string as an argument and outputs it in a graphic representation using ASCII characters. Each character is rendered with a height of 8 lines using one of three available banner styles.

## Features

- Support for letters, numbers, spaces, and special characters
- Three banner styles: `standard`, `shadow`, `thinkertoy`
- Handles newline characters (`\n`) for multi-line output
- 8-line character height format
- Standard Go packages only

## Project Structure

```
ascii-art/
├── cmd/
│   └── ascii-art/
│       └── main.go        # Entry point and CLI handling
├── internal/
│   ├── ascii/
│   │   ├── art.go         # Core ASCII art generation
│   │   ├── art_test.go    # Unit tests for art functions
│   │   ├── banner.go      # Banner file loading and parsing
│   │   └── banner_test.go # Unit tests for banner functions
│   └── version/
│       └── version.go     # Version information
├── assets/
│   ├── standard.txt       # Standard banner template
│   ├── shadow.txt         # Shadow banner template
│   └── thinkertoy.txt     # Thinkertoy banner template
├── .gitignore
├── go.mod
└── README.md
```

## Banner Format

- Each character has exactly 8 lines
- Characters are separated by newlines
- ASCII values 32-126 are supported
- Files contain 95 characters (space through tilde)

## Build Instructions

### Prerequisites

- Go 1.19 or higher
- Unix-like environment (Linux/macOS) or Windows with Go installed

### Setup

1. **Clone/Initialize the project:**
   ```bash
   git clone https://platform.zone01.gr/git/glalioti/ascii-art.git
   cd ascii-art
   ```

2. **Initialize Go module:**
   ```bash
   go mod init ascii-art
   ```

3. **Create directory structure:**
   ```bash
   mkdir -p cmd/ascii-art internal/ascii internal/version assets
   ```

### Building

1. **Build the executable:**
   ```bash
   go build -o ascii-art ./cmd/ascii-art
   ```

2. **Run directly with Go:**
   ```bash
   go run ./cmd/ascii-art "your text here"
   ```

### Testing

1. **Run all tests:**
   ```bash
   go test ./...
   ```

2. **Run tests with coverage:**
   ```bash
   go test -cover ./...
   ```

3. **Run specific package tests:**
   ```bash
   go test ./internal/ascii
   ```

### Usage Examples

```bash
# Basic usage
go run ./cmd/ascii-art "Hello"

# With newlines
go run ./cmd/ascii-art "Hello\nWorld"

# Empty string
go run ./cmd/ascii-art ""

# Special characters
go run ./cmd/ascii-art "Hello There!"

# Numbers and mixed content
go run ./cmd/ascii-art "123 ABC"
```

## Development Guidelines

### Code Organization

- **cmd/ascii-art/main.go**: Handle command-line arguments and program entry
- **internal/ascii/art.go**: Core logic for converting strings to ASCII art
- **internal/ascii/art_test.go**: Unit tests for ASCII generation
- **internal/ascii/banner.go**: Banner file loading and character mapping
- **internal/ascii/banner_test.go**: Unit tests for banner loading
- **internal/version/version.go**: Version and build information
- **assets/**: Banner template files

### Best Practices

- Use standard Go formatting (`go fmt`)
- Follow Go naming conventions
- Write comprehensive tests
- Handle edge cases (empty strings, invalid characters)
- Use proper error handling
- Keep functions focused and testable

### Error Handling

The program should handle:
- Invalid command-line arguments
- Missing banner files
- Unsupported characters
- File reading errors

### Performance Considerations

- Load banner files once at startup
- Use efficient string building for output
- Minimize memory allocations

## Implementation Phases

### Phase 1: Core Structure
- Set up project structure
- Create basic main.go with argument parsing
- Implement banner file loading

### Phase 2: ASCII Generation
- Implement character-to-ASCII conversion
- Handle basic string processing
- Add newline support

### Phase 3: Testing & Polish
- Write comprehensive tests
- Add error handling
- Optimize performance
- Documentation cleanup

## Testing Strategy

### Unit Tests
- Banner file loading
- Character mapping
- String processing
- Edge cases (empty strings, special chars)

### Integration Tests
- End-to-end ASCII generation
- Multiple banner styles
- Complex input strings

### Test Data
- Use provided banner files
- Test with various input combinations
- Validate output format consistency

## Deployment

### Local Development
```bash
go run ./cmd/ascii-art "test string"
```

### Production Build
```bash
go build -ldflags="-s -w" -o ascii-art ./cmd/ascii-art
./ascii-art "production text"
```

This documentation provides a clear roadmap for implementing your ASCII-art project with proper structure and testing approach.