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
	}{
		{
			name:    "empty string",
			input:   "",
			wantErr: false,
		},
		{
			name:    "single character",
			input:   "A",
			wantErr: false,
		},
		{
			name:    "simple word",
			input:   "Hello",
			wantErr: false,
		},
		{
			name:    "with newline",
			input:   "Hello\\nWorld",
			wantErr: false,
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

			// For multi-line input, check that we have the right structure
			if strings.Contains(tt.input, "\\n") {
				lines := strings.Split(result, "\n")
				if len(lines) < 8 {
					t.Errorf("Expected at least 8 lines for ASCII art, got %d", len(lines))
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