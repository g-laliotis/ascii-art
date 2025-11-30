package ascii

import (
	"strings"
	"testing"
)

func TestGenerateArt(t *testing.T) {
	// Load banner for testing
	charMap, err := LoadBanner("../../assets/standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	tests := []struct {
		name     string
		input    string
		wantErr  bool
		contains string
	}{
		{
			name:    "empty string",
			input:   "",
			wantErr: false,
			contains: "",
		},
		{
			name:    "single character A",
			input:   "A",
			wantErr: false,
			contains: "/\\",
		},
		{
			name:    "single character a",
			input:   "a",
			wantErr: false,
			contains: "__ _",
		},
		{
			name:    "simple word",
			input:   "Hello",
			wantErr: false,
			contains: "_    _",
		},
		{
			name:    "with newline",
			input:   "Hello\\nWorld",
			wantErr: false,
			contains: "$",
		},
		{
			name:    "numbers",
			input:   "123",
			wantErr: false,
			contains: "/ |",
		},
		{
			name:    "special characters",
			input:   "!@#",
			wantErr: false,
			contains: "|",
		},
		{
			name:    "fat underscore",
			input:   "_",
			wantErr: false,
			contains: "|_______|",
		},
		{
			name:    "space character",
			input:   " ",
			wantErr: false,
			contains: "$",
		},
		{
			name:    "mixed case",
			input:   "HeLLo",
			wantErr: false,
			contains: "_    _",
		},
		{
			name:    "all uppercase",
			input:   "HELLO",
			wantErr: false,
			contains: "_    _",
		},
		{
			name:    "all lowercase",
			input:   "hello",
			wantErr: false,
			contains: "_",
		},
		{
			name:    "double newline",
			input:   "Hi\\n\\nBye",
			wantErr: false,
			contains: "$",
		},
		{
			name:    "only newline",
			input:   "\\n",
			wantErr: false,
			contains: "$",
		},
		{
			name:    "punctuation",
			input:   ".,;:!?",
			wantErr: false,
			contains: "_",
		},
		{
			name:    "brackets and symbols",
			input:   "[]{}()",
			wantErr: false,
			contains: "|",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateArt(tt.input, charMap)
			
			if tt.input == "" {
				if result != "" {
					t.Errorf("Expected empty result for empty input, got: %s", result)
				}
				return
			}

			// Check that result has content
			if result == "" && tt.input != "" {
				t.Error("Expected non-empty result for non-empty input")
			}

			// Check for expected content
			if tt.contains != "" && !strings.Contains(result, tt.contains) {
				t.Errorf("Expected result to contain '%s', got: %s", tt.contains, result)
			}

			// For multi-line input, check that we have the right structure
			if strings.Contains(tt.input, "\\n") {
				lines := strings.Split(result, "\n")
				if len(lines) < 8 {
					t.Errorf("Expected at least 8 lines for ASCII art, got %d", len(lines))
				}
			}

			// Check that all lines end with $ (except empty lines from \n\n)
			lines := strings.Split(result, "\n")
			for i, line := range lines {
				if line != "" && line != "$" && !strings.HasSuffix(line, "$") {
					t.Errorf("Line %d should end with $, got: '%s'", i, line)
				}
			}
		})
	}
}

func TestGenerateLineArt(t *testing.T) {
	charMap, err := LoadBanner("../../assets/standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	result := generateLineArt("A", charMap)
	if len(result) != 8 {
		t.Errorf("Expected 8 lines for character art, got %d", len(result))
	}

	// Test empty line
	result = generateLineArt("", charMap)
	if len(result) != 1 || result[0] != "" {
		t.Error("Expected single empty string for empty line")
	}
}