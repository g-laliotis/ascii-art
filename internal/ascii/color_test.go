package ascii

import (
	"strings"
	"testing"
)

func TestApplyColor(t *testing.T) {
	charMap := map[rune][]string{
		'a': {"  __ _  ", " / _` | ", "| (_| | ", " \\__,_| ", "        ", "        ", "        ", "        "},
		'b': {" _      ", "| |     ", "| |__   ", "| '_ \\  ", "| |_) | ", "|_.__/  ", "        ", "        "},
	}

	tests := []struct {
		name      string
		artLines  []string
		substring string
		color     string
		text      string
		wantColor bool
	}{
		{
			name:      "Color entire output",
			artLines:  []string{"test$", "line$"},
			substring: "",
			color:     "red",
			text:      "ab",
			wantColor: true,
		},
		{
			name:      "Color substring",
			artLines:  []string{"  __ _  $"},
			substring: "a",
			color:     "blue",
			text:      "a",
			wantColor: true,
		},
		{
			name:      "Invalid color",
			artLines:  []string{"test$"},
			substring: "",
			color:     "invalid",
			text:      "ab",
			wantColor: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ApplyColor(tt.artLines, tt.substring, tt.color, tt.text, charMap)
			hasColor := strings.Contains(strings.Join(result, ""), "\033[")
			if hasColor != tt.wantColor {
				t.Errorf("ApplyColor() hasColor = %v, want %v", hasColor, tt.wantColor)
			}
		})
	}
}

func TestFindSubstringIndices(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		substring string
		want      []int
	}{
		{"Single occurrence", "hello", "ll", []int{2}},
		{"Multiple occurrences", "a king kitten have kit", "kit", []int{7, 19}},
		{"No occurrence", "hello", "xyz", []int{}},
		{"Overlapping", "aaa", "aa", []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSubstringIndices(tt.text, tt.substring)
			if len(got) != len(tt.want) {
				t.Errorf("findSubstringIndices() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("findSubstringIndices()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}
