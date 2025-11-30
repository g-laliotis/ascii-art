# ASCII-Art

> Transform your text into beautiful ASCII art

[![CI](https://github.com/g-laliotis/ascii-art/actions/workflows/ci.yml/badge.svg)](https://github.com/g-laliotis/ascii-art/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/g-laliotis/ascii-art)](https://goreportcard.com/report/github.com/g-laliotis/ascii-art)
[![Makefile](https://img.shields.io/badge/build-Makefile-blue.svg)](Makefile)

ASCII-Art is a command-line tool that converts regular text into stylized ASCII art using predefined banner templates. Perfect for creating eye-catching headers, banners, or just having fun with text!

🌐 **[Try the Live Demo](https://g-laliotis.github.io/ascii-art/)**

## ✨ Features

- 🎨 ASCII art using `standard` banner style (with `shadow` and `thinkertoy` support planned)
- 📝 Support for letters, numbers, spaces, and special characters
- 🔄 Multi-line output with `\n` support
- ⚡ Fast and lightweight - uses only Go standard library
- 🎯 Simple command-line interface

## 🚀 Quick Start

```bash
# Clone the repository
git clone https://github.com/g-laliotis/ascii-art.git
cd ascii-art

# Run directly
go run ./cmd/ascii-art "Hello World"

# Or use Makefile
make run
```

## 📝 Example Output

```bash
$ go run ./cmd/ascii-art "Hello"
$
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
```

## 🛠️ Installation

### Prerequisites
- Go 1.19 or higher

### Using Makefile

```bash
# Build executable
make build

# Run tests
make test

# Install to GOPATH/bin
make install

# See all available commands
make help
```

### Manual Build

```bash
# Build executable
go build -o ascii-art ./cmd/ascii-art

# Run all tests
go test -v ./...
```

## 📚 Usage

```bash
# Basic text
go run ./cmd/ascii-art "Hello"

# Multi-line text
go run ./cmd/ascii-art "Hello\nWorld"

# Special characters and numbers
go run ./cmd/ascii-art "Hello There! 123"

# Empty string (prints nothing)
go run ./cmd/ascii-art ""

# Long text (save to file for better viewing)
go run ./cmd/ascii-art "ABCDEFGHIJKLMNOPQRSTUVWXYZ" > output.txt
```

### 📏 Viewing Long Output

For very long strings (like the full alphabet), the output may be too wide for your terminal:

**Problem**: Lines wrap and appear distorted in terminal

**Solutions**:
1. **Save to file**: `go run ./cmd/ascii-art "long text" > output.txt`
2. **Use wider terminal**: Increase terminal width
3. **Text editor**: Open output file with word wrap disabled
4. **Horizontal scroll**: Use `less -S output.txt` for scrollable viewing

## 📁 Project Structure

```
ascii-art/
├── cmd/ascii-art/main.go      # Entry point
├── internal/ascii/            # Core logic and tests
├── internal/version/          # Version info
├── assets/
│   └── standard.txt           # Standard banner template
├── go.mod                     # Go module
└── LICENSE                    # MIT License
```

## 📖 How It Works

1. **Input**: Takes a string as command-line argument
2. **Processing**: Maps each character to its 8-line ASCII representation
3. **Output**: Combines characters horizontally to create ASCII art
4. **Format**: Each character is exactly 8 lines tall
5. **Support**: ASCII characters 32-126 (printable characters)
6. **Width**: Characters have variable widths (W is wider than I, etc.)

## 🤝 Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

This project follows Go best practices:
- Standard Go formatting (`go fmt`)
- Comprehensive unit tests (100% coverage)
- Clean, readable code structure
- Proper error handling
- CI/CD with GitHub Actions

## 👥 Authors

- **Giorgos Laliotis**
- **Stavros Gkraikas**

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🌟 Links

- 🌐 [Live Demo & Documentation](https://g-laliotis.github.io/ascii-art/)
- 📚 [Contributing Guidelines](CONTRIBUTING.md)
- 🐛 [Report Issues](https://github.com/g-laliotis/ascii-art/issues)
- 💡 [Feature Requests](https://github.com/g-laliotis/ascii-art/issues)

## 🎓 Acknowledgments

This project is part of the Zone01 curriculum and is now open-source for the community.

