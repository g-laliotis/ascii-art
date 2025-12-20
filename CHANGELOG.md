# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.2.0] - 2024-12-20

### Added
- **Multiple banner support**: Added `shadow` and `thinkertoy` banner styles alongside existing `standard`
- **File output functionality**: New `--output=filename` flag to save ASCII art to files
- **Cross-platform terminal width detection**: Separate implementations for Unix/Linux/macOS and Windows
- **Enhanced argument parsing**: Support for banner selection as command argument
- **Comprehensive test suite**: Added banner-specific tests and output functionality tests
- **OS-specific modules**: `terminal_unix.go` and `terminal_windows.go` for platform compatibility
- **File output module**: New `output.go` with `SaveToFile` function
- **Enhanced Makefile**: Added targets for different banners, file output, and combined features
- **Improved documentation**: Updated README, CONTRIBUTING, and project structure

### Changed
- **Terminal width detection**: Refactored from single function to OS-specific implementations
- **Color application logic**: Improved substring coloring to handle multiple occurrences correctly
- **Default terminal width**: Changed from 80 to 200 characters for better display
- **Project structure**: Added new files for cross-platform support and enhanced testing
- **Command-line interface**: Enhanced to support banner selection and file output

### Fixed
- **Color substring handling**: Fixed multiple occurrence coloring in same text
- **Cross-platform compatibility**: Proper terminal width detection on Windows and Unix systems
- **Banner loading**: Improved error handling for different banner files

### Technical Details
- Added build tags for OS-specific code (`//go:build unix`, `//go:build windows`)
- Enhanced test coverage with banner-specific and output functionality tests
- Improved code organization with separate modules for different functionalities
- Maintained backward compatibility while adding new features

## [1.1.0] - 2024-12-20

### Added
- Color support with `--color=<color>` flag
- ANSI color codes: red, green, yellow, blue, magenta, cyan, white, orange
- Ability to color entire ASCII art output
- Ability to color specific substrings within output
- Support for multiple occurrences of substring coloring
- Accurate character position calculation for substring coloring
- Usage message for incorrect flag format
- Comprehensive unit tests for color functionality
- Updated README with color examples and file structure

## [1.0.0] - 2024-12-19

### Added
- Initial release of ASCII-Art generator
- Support for standard banner template with all printable ASCII characters (32-126)
- Multi-line text support with `\n` handling
- Command-line interface with single argument input
- Comprehensive test suite with 100% coverage (50+ test cases)
- Professional project structure with `cmd/` and `internal/` directories
- Makefile with build, test, install, and clean targets
- MIT license and open-source ready documentation
- GitHub Actions CI/CD pipeline for automated testing
- Contributing guidelines and code of conduct
- GitHub Pages website with live demo
- Support for letters, numbers, spaces, and special characters

### Fixed
- Banner parsing logic to correctly read all 8 lines for each character
- Character 'g' descender display issue - now shows proper `|___/` bottom line
- Fat underscore character alignment (9 characters width)
- Multi-line output formatting with proper `$` line endings
- Empty line handling for `\n\n` sequences
- Go module version compatibility (changed from invalid 1.25.0 to 1.21)

### Technical Details
- Built with Go 1.21+ using only standard library
- Zero external dependencies for maximum portability
- Each ASCII character rendered as exactly 8 lines tall
- Variable character widths for proper proportions
- Efficient character mapping and horizontal combination
- Proper error handling and input validation

### Documentation
- Complete README with installation, usage, and examples
- Live demo website at https://g-laliotis.github.io/ascii-art/
- Contributing guidelines for open-source development
- Professional badges (CI, license, Go report card, Makefile)
- Comprehensive usage examples and troubleshooting guide

[Unreleased]: https://github.com/g-laliotis/ascii-art/compare/v1.2.0...HEAD
[1.2.0]: https://github.com/g-laliotis/ascii-art/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/g-laliotis/ascii-art/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/g-laliotis/ascii-art/releases/tag/v1.0.0