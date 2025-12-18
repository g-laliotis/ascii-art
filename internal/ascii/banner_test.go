package ascii

import (
	"testing"
)

func TestLoadBanner(t *testing.T) {
	// Test loading the standard banner file
	charMap, err := LoadBanner("../../assets/standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	// Test that we have the expected number of characters
	if len(charMap) == 0 {
		t.Error("Character map is empty")
	}

	// Test that space character (32) exists
	if _, exists := charMap[' ']; !exists {
		t.Error("Space character not found in banner")
	}

	// Test that 'A' character (65) exists and has 8 lines
	if lines, exists := charMap['A']; !exists {
		t.Error("Character 'A' not found in banner")
	} else if len(lines) != 8 {
		t.Errorf("Character 'A' should have 8 lines, got %d", len(lines))
	}
}

func TestLoadBannerFileNotFound(t *testing.T) {
	_, err := LoadBanner("nonexistent.txt")
	if err == nil {
		t.Error("Expected error for nonexistent file")
	}
}

func TestLoadMultipleBanners(t *testing.T) {
	banners := []string{"../../assets/standard.txt", "../../assets/thinkertoy.txt"}
	testChars := []rune{'A', 'a', '1', '!', ' '}

	for _, banner := range banners {
		t.Run(banner, func(t *testing.T) {
			charMap, err := LoadBanner(banner)
			if err != nil {
				t.Errorf("LoadBanner(%s) error = %v", banner, err)
				return
			}

			for _, char := range testChars {
				if lines, exists := charMap[char]; !exists {
					t.Errorf("Banner %s missing character %c", banner, char)
				} else if len(lines) != 8 {
					t.Errorf("Banner %s character %c has %d lines, want 8", banner, char, len(lines))
				}
			}
		})
	}
}