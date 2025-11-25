# ASCII-Art Architecture

## System Design

### High-Level Architecture

```
Input String → Argument Parser → ASCII Generator → Output
                     ↓
               Banner Loader ← Banner Files
```

### Component Breakdown

#### 1. Main Entry Point (`main.go`)
- **Responsibility**: Command-line argument handling and program orchestration
- **Functions**:
  - Parse command-line arguments
  - Validate input
  - Initialize ASCII generator
  - Handle program exit codes

#### 2. Banner Management (`ascii/banner.go`)
- **Responsibility**: Loading and parsing banner template files
- **Key Functions**:
  - `LoadBanner(filename string) (map[rune][]string, error)`
  - `ParseBannerFile(content string) map[rune][]string`
- **Data Structure**: Map of rune to 8-line string slice

#### 3. ASCII Art Generator (`ascii/art.go`)
- **Responsibility**: Convert input strings to ASCII art
- **Key Functions**:
  - `GenerateArt(text string, banner map[rune][]string) string`
  - `ProcessLine(line string, banner map[rune][]string) []string`
  - `CombineCharacters(chars [][]string) string`

### Data Flow

1. **Input Processing**:
   ```
   Command Line Args → String Validation → Character Parsing
   ```

2. **Banner Loading**:
   ```
   Banner File → File Reading → Character Mapping → Memory Cache
   ```

3. **Art Generation**:
   ```
   Input String → Line Splitting → Character Lookup → Line Assembly → Output
   ```

### Key Algorithms

#### Character Mapping Algorithm
```
For each character in ASCII 32-126:
  - Read 8 consecutive lines from banner file
  - Map character code to line array
  - Store in lookup table
```

#### Art Generation Algorithm
```
For each line in input:
  For each character in line:
    - Lookup character in banner map
    - Collect 8-line representation
  - Combine all characters horizontally
  - Output combined lines
```

### Error Handling Strategy

#### Input Validation
- Empty arguments
- Invalid characters (outside ASCII 32-126)
- Malformed newline sequences

#### File Operations
- Missing banner files
- Corrupted banner format
- File permission issues

#### Runtime Errors
- Memory allocation failures
- Invalid character mappings

### Performance Considerations

#### Memory Management
- Load banners once at startup
- Use string builder for output construction
- Minimize string concatenation

#### Time Complexity
- O(n) for input processing (n = input length)
- O(1) for character lookup
- O(n*8) for output generation

### Testing Architecture

#### Unit Test Structure
```
tests/
├── banner_test.go     # Banner loading tests
├── art_test.go        # ASCII generation tests
└── integration_test.go # End-to-end tests
```

#### Test Categories
1. **Unit Tests**: Individual function testing
2. **Integration Tests**: Component interaction testing
3. **Edge Case Tests**: Boundary condition testing
4. **Performance Tests**: Load and stress testing

### Configuration Management

#### Banner File Format
- Fixed 8-line height per character
- 95 characters (ASCII 32-126)
- Newline-separated character definitions

#### Project Configuration
- Go module configuration
- Build flags and optimization
- Test configuration

This architecture ensures maintainable, testable, and efficient ASCII art generation.