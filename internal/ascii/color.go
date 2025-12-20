package ascii

import "strings"

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

	// Apply color to all occurrences at once
	return colorAllSubstrings(artLines, indices, len(substring), colorCode, originalText, charMap)
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

// colorAllSubstrings applies color to all substring occurrences at once
func colorAllSubstrings(artLines []string, indices []int, length int, colorCode, originalText string, charMap map[rune][]string) []string {
	// Calculate pixel positions for all occurrences
	type colorRange struct {
		start, end int
	}
	var ranges []colorRange

	for _, startIdx := range indices {
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

		ranges = append(ranges, colorRange{pixelStart, pixelEnd})
	}

	// Apply color to each line
	for i := range artLines {
		if artLines[i] == "" || artLines[i] == "$" {
			continue
		}

		line := strings.TrimSuffix(artLines[i], "$")
		result := ""
		lastPos := 0

		for _, r := range ranges {
			if r.start < len(line) && r.end <= len(line) {
				// Add text before colored section
				if r.start > lastPos {
					result += line[lastPos:r.start]
				}
				// Add colored section
				result += colorCode + line[r.start:r.end] + colorMap["reset"]
				lastPos = r.end
			}
		}

		// Add remaining text
		if lastPos < len(line) {
			result += line[lastPos:]
		}

		artLines[i] = result + "$"
	}

	return artLines
}
