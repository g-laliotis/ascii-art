# ASCII-Art

> Transform your text into beautiful ASCII art

ASCII-Art is a command-line tool that converts regular text into stylized ASCII art using predefined banner templates. Perfect for creating eye-catching headers, banners, or just having fun with text!

## ✨ Features

- 🎨 ASCII art using `standard` banner style (with `shadow` and `thinkertoy` support planned)
- 📝 Support for letters, numbers, spaces, and special characters
- 🔄 Multi-line output with `\n` support
- ⚡ Fast and lightweight - uses only Go standard library
- 🎯 Simple command-line interface

## 🚀 Quick Start

```bash
# Clone the repository
git clone https://platform.zone01.gr/git/glalioti/ascii-art.git
cd ascii-art

# Initialize Go module
go mod init ascii-art

# Run directly
go run . "Hello World"
```

## 📝 Example Output

```bash
$ go run . "Hello"
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
```

## 🛠️ Installation

### Prerequisites
- Go 1.19 or higher

### Build from Source

```bash
# Build executable
go build -o ascii-art .

# Run the binary
./ascii-art "Your Text Here"
```

### Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...
```

## 📚 Usage

```bash
# Basic text
go run . "Hello"

# Multi-line text
go run . "Hello\nWorld"

# Special characters and numbers
go run . "Hello There! 123"

# Empty string (prints nothing)
go run . ""
```

## 📁 Project Structure

```
ascii-art/
├── main.go                    # Entry point
├── cmd/ascii-art/main.go      # Alternative entry point
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

## 🤝 Contributing

This project follows Go best practices:
- Standard Go formatting (`go fmt`)
- Comprehensive unit tests
- Clean, readable code structure
- Proper error handling

## 👥 Authors

- **Giorgos Laliotis**
- **Stavros Gkraikas**

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🎓 Acknowledgments

This project is part of the Zone01 curriculum.

