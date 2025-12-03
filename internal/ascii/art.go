package ascii

import (
	"os"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
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

// getTerminalWidth returns the terminal width, defaults to 80 if unable to detect
func getTerminalWidth() int {
	type winsize struct {
		Row    uint16
		Col    uint16
		Xpixel uint16
		Ypixel uint16
	}

	ws := &winsize{}
	retVal, _, _ := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(os.Stdout.Fd()),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retVal) == -1 {
		// Fallback: try COLUMNS environment variable
		if cols := os.Getenv("COLUMNS"); cols != "" {
			if width, err := strconv.Atoi(cols); err == nil {
				return width
			}
		}
		// Default fallback
		return 80
	}
	return int(ws.Col)
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