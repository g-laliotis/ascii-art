package ascii

import (
	"strings"
)

// GenerateArt converts input text to ASCII art using the provided character map
func GenerateArt(text string, charMap map[rune][]string) string {
	if text == "" {
		return ""
	}

	// Split text by newlines to handle multi-line input
	lines := strings.Split(text, "\\n")
	var result []string

	for _, line := range lines {
		if line == "" {
			// Add empty line with $
			result = append(result, "$")
			continue
		}

		// Generate ASCII art for this line
		artLines := generateLineArt(line, charMap)
		result = append(result, artLines...)
	}

	return strings.Join(result, "\n")
}

// generateLineArt converts a single line of text to ASCII art
func generateLineArt(line string, charMap map[rune][]string) []string {
	if line == "" {
		return []string{""}
	}

	// Initialize 8 lines for the ASCII art
	artLines := make([]string, 8)

	// Process each character in the line
	for _, char := range line {
		if charLines, exists := charMap[char]; exists {
			// Add each line of the character to the corresponding art line
			for i := 0; i < 8; i++ {
				if i < len(charLines) {
					artLines[i] += charLines[i]
				}
			}
		}
	}

	// Add $ at the end of each non-empty line
	for i := range artLines {
		if artLines[i] != "" {
			artLines[i] += "$"
		}
	}

	return artLines
}