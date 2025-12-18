package ascii

import (
	"strings"
	"testing"
)

func TestGenerateArtWithDifferentBanners(t *testing.T) {
	tests := []struct {
		name   string
		banner string
		text   string
	}{
		{"Standard Hello", "../../assets/standard.txt", "Hello"},
		{"Thinkertoy Hello", "../../assets/thinkertoy.txt", "Hello"},
		{"Standard Numbers", "../../assets/standard.txt", "123"},
		{"Thinkertoy Numbers", "../../assets/thinkertoy.txt", "123"},
		{"Standard Special", "../../assets/standard.txt", "A!B"},
		{"Thinkertoy Special", "../../assets/thinkertoy.txt", "A!B"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			charMap, err := LoadBanner(tt.banner)
			if err != nil {
				t.Errorf("LoadBanner() error = %v", err)
				return
			}

			result := GenerateArt(tt.text, charMap)
			if result == "" {
				t.Error("GenerateArt() returned empty result")
				return
			}

			lines := strings.Split(result, "\n")
			if len(lines) < 8 {
				t.Errorf("GenerateArt() returned %d lines, want at least 8", len(lines))
			}

			// Check that all lines end with $
			for i, line := range lines {
				if line != "" && !strings.HasSuffix(line, "$") {
					t.Errorf("Line %d does not end with $: %q", i, line)
				}
			}
		})
	}
}

func TestBannerConsistency(t *testing.T) {
	banners := []string{"../../assets/standard.txt", "../../assets/thinkertoy.txt"}
	
	for _, banner := range banners {
		t.Run(banner, func(t *testing.T) {
			charMap, err := LoadBanner(banner)
			if err != nil {
				t.Errorf("LoadBanner() error = %v", err)
				return
			}

			// Test that all ASCII printable characters are present
			for i := 32; i <= 126; i++ {
				char := rune(i)
				if lines, exists := charMap[char]; !exists {
					t.Errorf("Missing character %c (ASCII %d)", char, i)
				} else if len(lines) != 8 {
					t.Errorf("Character %c has %d lines, want 8", char, len(lines))
				}
			}
		})
	}
}