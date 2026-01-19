package ascii

import (
	"strings"
	"testing"
)

func TestApplyAlignment(t *testing.T) {
	// Sample ASCII art lines for testing
	testLines := []string{
		" _    _          _   _          $",
		"| |  | |        | | | |         $",
		"| |__| |   ___  | | | |   ___   $",
		"|  __  |  / _ \\ | | | |  / _ \\  $",
		"| |  | | |  __/ | | | | | (_) | $",
		"|_|  |_|  \\___| |_| |_|  \\___/  $",
		"                                $",
		"                                $",
	}

	tests := []struct {
		name      string
		alignment string
		lines     []string
		wantLen   int
	}{
		{
			name:      "Left alignment (default)",
			alignment: "left",
			lines:     testLines,
			wantLen:   len(testLines),
		},
		{
			name:      "Empty alignment (default to left)",
			alignment: "",
			lines:     testLines,
			wantLen:   len(testLines),
		},
		{
			name:      "Right alignment",
			alignment: "right",
			lines:     testLines,
			wantLen:   len(testLines),
		},
		{
			name:      "Center alignment",
			alignment: "center",
			lines:     testLines,
			wantLen:   len(testLines),
		},
		{
			name:      "Justify alignment",
			alignment: "justify",
			lines:     testLines,
			wantLen:   len(testLines),
		},
		{
			name:      "Invalid alignment (fallback to default)",
			alignment: "invalid",
			lines:     testLines,
			wantLen:   len(testLines),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ApplyAlignment(tt.lines, tt.alignment)
			
			// Check that we get the same number of lines
			if len(result) != tt.wantLen {
				t.Errorf("ApplyAlignment() returned %d lines, want %d", len(result), tt.wantLen)
			}
			
			// Check that all lines end with $
			for i, line := range result {
				if line != "" && line != "$" && !strings.HasSuffix(line, "$") {
					t.Errorf("Line %d does not end with $: %q", i, line)
				}
			}
		})
	}
}

func TestAlignRight(t *testing.T) {
	tests := []struct {
		name      string
		lines     []string
		termWidth int
		want      []string
	}{
		{
			name:      "Simple right alignment",
			lines:     []string{"hello$", "world$"},
			termWidth: 20,
			want:      []string{"              hello$", "              world$"},
		},
		{
			name:      "Empty line handling",
			lines:     []string{"hello$", "$", "world$"},
			termWidth: 20,
			want:      []string{"              hello$", "$", "              world$"},
		},
		{
			name:      "Content too wide",
			lines:     []string{"very long content that exceeds terminal width$"},
			termWidth: 10,
			want:      []string{"very long content that exceeds terminal width$"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := alignRightConsistent(tt.lines, tt.termWidth)
			
			if len(result) != len(tt.want) {
				t.Errorf("alignRight() returned %d lines, want %d", len(result), len(tt.want))
				return
			}
			
			for i, line := range result {
				if line != tt.want[i] {
					t.Errorf("alignRight() line %d = %q, want %q", i, line, tt.want[i])
				}
			}
		})
	}
}

func TestAlignCenter(t *testing.T) {
	tests := []struct {
		name      string
		lines     []string
		termWidth int
		want      []string
	}{
		{
			name:      "Simple center alignment",
			lines:     []string{"hello$"},
			termWidth: 20,
			want:      []string{"       hello$"},
		},
		{
			name:      "Even width centering",
			lines:     []string{"test$"},
			termWidth: 10,
			want:      []string{"  test$"},
		},
		{
			name:      "Odd width centering",
			lines:     []string{"test$"},
			termWidth: 11,
			want:      []string{"   test$"},
		},
		{
			name:      "Empty line handling",
			lines:     []string{"$"},
			termWidth: 20,
			want:      []string{"$"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := alignCenterConsistent(tt.lines, tt.termWidth)
			
			if len(result) != len(tt.want) {
				t.Errorf("alignCenter() returned %d lines, want %d", len(result), len(tt.want))
				return
			}
			
			for i, line := range result {
				if line != tt.want[i] {
					t.Errorf("alignCenter() line %d = %q, want %q", i, line, tt.want[i])
				}
			}
		})
	}
}

func TestAlignJustify(t *testing.T) {
	tests := []struct {
		name      string
		lines     []string
		termWidth int
	}{
		{
			name:      "Simple justify alignment",
			lines:     []string{"hello$"},
			termWidth: 20,
		},
		{
			name:      "Multiple lines justify",
			lines:     []string{"hello$", "world$"},
			termWidth: 30,
		},
		{
			name:      "Empty line handling",
			lines:     []string{"hello$", "$", "world$"},
			termWidth: 25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := alignJustifyConsistent(tt.lines, tt.termWidth, "")
			
			// Check that we get the same number of lines
			if len(result) != len(tt.lines) {
				t.Errorf("alignJustify() returned %d lines, want %d", len(result), len(tt.lines))
			}
			
			// Check that all non-empty lines end with $
			for i, line := range result {
				if line != "" && line != "$" && !strings.HasSuffix(line, "$") {
					t.Errorf("Line %d does not end with $: %q", i, line)
				}
			}
			
			// Check that content is distributed (not same as original for non-empty lines)
			for i := range result {
				if tt.lines[i] != "$" && tt.lines[i] != "" {
					// For justify with empty originalText, it should be left-aligned (no change)
					// So we just check it doesn't crash and returns same number of lines
				}
			}
		})
	}
}

func TestAlignmentWithColors(t *testing.T) {
	// Test that alignment works with ANSI color codes
	coloredLines := []string{
		"\033[31mhello\033[0m$",
		"\033[32mworld\033[0m$",
	}
	
	tests := []struct {
		name      string
		alignment string
		lines     []string
	}{
		{"Right with colors", "right", coloredLines},
		{"Center with colors", "center", coloredLines},
		{"Justify with colors", "justify", coloredLines},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ApplyAlignment(tt.lines, tt.alignment)
			
			// Check that we still have the same number of lines
			if len(result) != len(tt.lines) {
				t.Errorf("ApplyAlignment() with colors returned %d lines, want %d", len(result), len(tt.lines))
			}
			
			// Check that color codes are preserved
			for _, line := range result {
				if line != "$" && line != "" {
					// Check if any color codes are preserved
					hasColorCodes := strings.Contains(line, "\033[") 
					originalHasColors := false
					for _, origLine := range tt.lines {
						if strings.Contains(origLine, "\033[") {
							originalHasColors = true
							break
						}
					}
					
					if originalHasColors && !hasColorCodes {
						t.Errorf("Color codes not preserved in alignment: %q", line)
					}
				}
			}
		})
	}
}

func TestAlignmentEdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		lines     []string
		alignment string
		termWidth int
	}{
		{
			name:      "Very narrow terminal",
			lines:     []string{"hello$"},
			alignment: "center",
			termWidth: 3,
		},
		{
			name:      "Empty input",
			lines:     []string{},
			alignment: "right",
			termWidth: 80,
		},
		{
			name:      "Only empty lines",
			lines:     []string{"$", "$", "$"},
			alignment: "center",
			termWidth: 50,
		},
		{
			name:      "Mixed empty and content lines",
			lines:     []string{"hello$", "$", "world$", "$"},
			alignment: "right",
			termWidth: 30,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Should not panic and should return same number of lines
			result := ApplyAlignment(tt.lines, tt.alignment)
			
			if len(result) != len(tt.lines) {
				t.Errorf("ApplyAlignment() edge case returned %d lines, want %d", len(result), len(tt.lines))
			}
		})
	}
}