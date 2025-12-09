package ascii

import "strings"

// ANSI color codes
var colorMap = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	"orange":  "\033[38;5;208m",
	"reset":   "\033[0m",
}

// ApplyColor applies color to specific substring in ASCII art
func ApplyColor(artLines []string, substring, color, originalText string, charMap map[rune][]string) []string {
	colorCode, exists := colorMap[strings.ToLower(color)]
	if !exists {
		return artLines
	}

	if substring == "" {
		// Color entire output
		for i := range artLines {
			if artLines[i] != "" && artLines[i] != "$" {
				line := strings.TrimSuffix(artLines[i], "$")
				artLines[i] = colorCode + line + colorMap["reset"] + "$"
			}
		}
		return artLines
	}

	// Find all occurrences of substring in original text
	indices := findSubstringIndices(originalText, substring)
	if len(indices) == 0 {
		return artLines
	}

	// Apply color to each occurrence
	for _, idx := range indices {
		artLines = colorSubstringInArt(artLines, idx, len(substring), colorCode, originalText, charMap)
	}

	return artLines
}

// findSubstringIndices finds all starting indices of substring in text
func findSubstringIndices(text, substring string) []int {
	var indices []int
	start := 0
	for {
		idx := strings.Index(text[start:], substring)
		if idx == -1 {
			break
		}
		indices = append(indices, start+idx)
		start += idx + 1
	}
	return indices
}

// colorSubstringInArt applies color to a substring at given position in art lines
func colorSubstringInArt(artLines []string, startIdx, length int, colorCode, originalText string, charMap map[rune][]string) []string {
	// Calculate pixel position in art based on character positions
	pixelStart := 0
	for i := 0; i < startIdx; i++ {
		if i < len(originalText) {
			if charLines, exists := charMap[rune(originalText[i])]; exists && len(charLines) > 0 {
				pixelStart += len(charLines[0])
			}
		}
	}

	pixelEnd := pixelStart
	for i := startIdx; i < startIdx+length && i < len(originalText); i++ {
		if charLines, exists := charMap[rune(originalText[i])]; exists && len(charLines) > 0 {
			pixelEnd += len(charLines[0])
		}
	}

	// Apply color to each line
	for i := range artLines {
		if artLines[i] == "" || artLines[i] == "$" {
			continue
		}

		line := strings.TrimSuffix(artLines[i], "$")
		if pixelStart < len(line) && pixelEnd <= len(line) {
			before := line[:pixelStart]
			middle := line[pixelStart:pixelEnd]
			after := line[pixelEnd:]
			artLines[i] = before + colorCode + middle + colorMap["reset"] + after + "$"
		}
	}

	return artLines
}
