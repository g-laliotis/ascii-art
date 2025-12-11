package ascii

import (
	"os"
	"strconv"
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

		// Generate ASCII art for this line with wrapping
		artLines := generateLineArtWithWrap(line, charMap)
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

// getTerminalWidth returns the terminal width, defaults to 200 if unable to detect
func getTerminalWidth() int {
	// Try OS-specific detection first
	if width := getTerminalWidthOS(); width > 0 {
		return width
	}
	
	// Try COLUMNS environment variable
	if cols := os.Getenv("COLUMNS"); cols != "" {
		if width, err := strconv.Atoi(cols); err == nil {
			return width
		}
	}
	
	// Default fallback
	return 200
}

// generateLineArtWithWrap generates ASCII art for a line with terminal width wrapping
func generateLineArtWithWrap(text string, charMap map[rune][]string) []string {
	termWidth := getTerminalWidth()
	// Reserve 2 characters for $ signs
	maxWidth := termWidth - 2
	
	if maxWidth < 10 {
		// Terminal too narrow, use original method
		return generateLineArt(text, charMap)
	}

	var allLines []string
	currentText := ""
	currentWidth := 0
	
	for _, char := range text {
		// Get character width
		charLines, exists := charMap[char]
		if !exists {
			continue
		}
		
		charWidth := 0
		if len(charLines) > 0 {
			charWidth = len(charLines[0])
		}
		
		// Check if adding this character would exceed terminal width
		if currentWidth+charWidth > maxWidth && currentText != "" {
			// Generate art for current text and add to result
			currentArt := generateLineArt(currentText, charMap)
			allLines = append(allLines, currentArt...)
			
			// Start new line with current character
			currentText = string(char)
			currentWidth = charWidth
		} else {
			// Add character to current line
			currentText += string(char)
			currentWidth += charWidth
		}
	}
	
	// Generate art for remaining text
	if currentText != "" {
		currentArt := generateLineArt(currentText, charMap)
		allLines = append(allLines, currentArt...)
	}
	
	return allLines
}