# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/g-laliotis/ascii-art/compare/v1.1.0...HEAD
[1.1.0]: https://github.com/g-laliotis/ascii-art/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/g-laliotis/ascii-art/releases/tag/v1.0.0