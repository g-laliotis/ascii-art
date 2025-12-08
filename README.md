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
- 📱 **Automatic terminal width detection and wrapping** - adapts to any screen size
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

# Long text (automatically wraps to terminal width)
go run ./cmd/ascii-art "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
```

### 📱 Terminal Width Adaptation

The program automatically detects your terminal width and wraps long text accordingly:

- **Smart wrapping**: Automatically breaks long text at character boundaries
- **Any terminal size**: Works on narrow mobile terminals to wide desktop screens
- **Preserves formatting**: Each wrapped section maintains proper ASCII art structure
- **Fallback support**: Uses COLUMNS environment variable or defaults to 80 characters

## 📁 Project Structure

```
ascii-art/
├── cmd/ascii-art/main.go      # Entry point
├── internal/
│   ├── ascii/                 # Core ASCII generation logic
│   │   ├── art.go            # ASCII art generation functions
│   │   ├── banner.go         # Banner file loading and parsing
│   │   ├── art_test.go       # Unit tests for art generation
│   │   └── banner_test.go    # Unit tests for banner loading
│   └── version/
│       └── version.go        # Version information
├── assets/
│   └── standard.txt          # Standard banner template (8 lines per character)
├── docs/
│   └── index.html           # GitHub Pages website with live demo
├── .github/
│   └── workflows/
│       └── ci.yml           # GitHub Actions CI/CD pipeline
├── main_test.go             # Integration tests using exec.Command
├── edge_cases_test.go       # Comprehensive edge case tests
├── go.mod                   # Go module definition
├── Makefile                 # Build automation (build, test, install)
├── README.md                # Project documentation
├── CHANGELOG.md             # Version history and changes
├── CONTRIBUTING.md          # Contribution guidelines
└── LICENSE                  # MIT License
```

## 📖 How It Works

1. **Input**: Takes a string as command-line argument
2. **Terminal Detection**: Automatically detects terminal width using system calls
3. **Processing**: Maps each character to its 8-line ASCII representation
4. **Smart Wrapping**: Calculates character widths and wraps when exceeding terminal width
5. **Output**: Combines characters horizontally with automatic line breaks
6. **Format**: Each character is exactly 8 lines tall
7. **Support**: ASCII characters 32-126 (printable characters)
8. **Adaptive Width**: Characters have variable widths, automatically handled

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

## 📋 Changelog

See [CHANGELOG.md](CHANGELOG.md) for a detailed history of changes, new features, and bug fixes.

## 🌟 Links

- 🌐 [Live Demo & Documentation](https://g-laliotis.github.io/ascii-art/)
- 📋 [Changelog](CHANGELOG.md)
- 📚 [Contributing Guidelines](CONTRIBUTING.md)
- 🐛 [Report Issues](https://github.com/g-laliotis/ascii-art/issues)
- 💡 [Feature Requests](https://github.com/g-laliotis/ascii-art/issues)

## 🎓 Acknowledgments

This project is part of the Zone01 curriculum and is now open-source for the community.

