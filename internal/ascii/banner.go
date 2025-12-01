package ascii

import (
	"bufio"
	"fmt"
	"os"
)

// LoadBanner loads a banner file and returns a map of characters to their ASCII representations
func LoadBanner(filename string) (map[rune][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open banner file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	
	// Read all lines from the file
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read banner file: %w", err)
	}

	return parseBannerLines(lines), nil
}

// parseBannerLines converts the banner file lines into a character map
func parseBannerLines(lines []string) map[rune][]string {
	charMap := make(map[rune][]string)
	
	// ASCII printable characters start from 32 (space) to 126 (~)
	char := rune(32)
	
	for i := 0; i < len(lines); i += 9 { // Each character takes 8 lines + 1 separator
		if i+8 >= len(lines) {
			break
		}
		
		// Extract 8 lines for this character (skip the first empty separator line)
		charLines := make([]string, 8)
		for j := 0; j < 8; j++ {
			if i+j+1 < len(lines) {
				charLines[j] = lines[i+j+1]
			}
		}
		
		// Special handling for underscore to make it fat like dash
		if char == 95 { // underscore
			// Make underscore fat by adding thickness
			modifiedLines := make([]string, 8)
			copy(modifiedLines, charLines)
			modifiedLines[6] = " _______ " // Add line above (9 chars)
			modifiedLines[7] = "|_______|" // Make it thick like dash (9 chars)
			charMap[char] = modifiedLines
		} else {
			charMap[char] = charLines
		}
		char++
		
		if char > 126 {
			break
		}
	}
	
	return charMap
}