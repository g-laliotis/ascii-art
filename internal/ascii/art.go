package ascii

import (
	"os"
	"strconv"
	"strings"
)

// ANSI color codes (shared with color.go)
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

// GenerateArt converts input text to ASCII art using the provided character map
func GenerateArt(text string, charMap map[rune][]string) string {
	return GenerateArtWithColor(text, charMap, "", "")
}

// GenerateArtWithColorAndAlignment converts input text to ASCII art with optional color and alignment support
func GenerateArtWithColorAndAlignment(text string, charMap map[rune][]string, substring, color, alignment string) string {
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

		// For justify alignment with multiple words, handle specially
		if alignment == "justify" && len(strings.Fields(line)) > 1 {
			artLines := generateJustifiedArt(line, charMap, substring, color)
			result = append(result, artLines...)
		} else {
			// Generate ASCII art for this line with wrapping, color, and alignment
			artLines := generateLineArtWithWrapColorAndAlignment(line, charMap, substring, color, alignment)
			result = append(result, artLines...)
		}
	}

	return strings.Join(result, "\n")
}

// generateJustifiedArt generates justified ASCII art by distributing words across terminal width
func generateJustifiedArt(text string, charMap map[rune][]string, substring, color string) []string {
	words := strings.Fields(text)
	if len(words) <= 1 {
		// Single word, use left alignment
		return generateSegmentWithColor(text, charMap, []struct{ start, end int }{}, 0, color)
	}

	termWidth := getTerminalWidth()
	maxWidth := termWidth - 1 // Reserve space for $
	
	// Generate ASCII art for each word and calculate widths
	var wordArts [][]string
	var wordWidths []int
	
	for _, word := range words {
		wordArt := generateSegmentWithColor(word, charMap, []struct{ start, end int }{}, 0, "")
		wordArts = append(wordArts, wordArt)
		
		// Calculate word width (from first line, excluding $)
		if len(wordArt) > 0 {
			wordWidth := len(strings.TrimSuffix(wordArt[0], "$"))
			wordWidths = append(wordWidths, wordWidth)
		}
	}
	
	// Group words into lines that fit within terminal width
	var lines [][]int // Each element contains word indices for that line
	currentLine := []int{}
	currentWidth := 0
	
	for i, width := range wordWidths {
		// Check if adding this word would exceed terminal width
		// Account for minimum 1 space between words
		requiredWidth := width
		if len(currentLine) > 0 {
			requiredWidth += 1 // Space before word
		}
		
		if currentWidth+requiredWidth > maxWidth && len(currentLine) > 0 {
			// Start new line
			lines = append(lines, currentLine)
			currentLine = []int{i}
			currentWidth = width
		} else {
			// Add to current line
			currentLine = append(currentLine, i)
			currentWidth += requiredWidth
		}
	}
	
	if len(currentLine) > 0 {
		lines = append(lines, currentLine)
	}
	
	// Generate justified output for each line
	var result []string
	
	for _, lineWords := range lines {
		if len(lineWords) == 1 {
			// Single word - left align
			wordArt := wordArts[lineWords[0]]
			result = append(result, wordArt...)
		} else {
			// Multiple words - justify
			totalWordWidth := 0
			for _, wordIdx := range lineWords {
				totalWordWidth += wordWidths[wordIdx]
			}
			
			// Calculate spacing
			totalSpacing := maxWidth - totalWordWidth
			gaps := len(lineWords) - 1
			spacePerGap := 1 // Minimum 1 space
			extraSpaces := 0
			
			if gaps > 0 && totalSpacing > gaps {
				spacePerGap = totalSpacing / gaps
				extraSpaces = totalSpacing % gaps
			}
			
			// Build justified lines
			lineResult := make([]string, 8)
			for lineIdx := 0; lineIdx < 8; lineIdx++ {
				for i, wordIdx := range lineWords {
					wordArt := wordArts[wordIdx]
					if lineIdx < len(wordArt) {
						content := strings.TrimSuffix(wordArt[lineIdx], "$")
						lineResult[lineIdx] += content
						
						// Add spacing after word (except last word)
						if i < len(lineWords)-1 {
							spaces := spacePerGap
							if i < extraSpaces {
								spaces++
							}
							lineResult[lineIdx] += strings.Repeat(" ", spaces)
						}
					}
				}
				lineResult[lineIdx] += "$"
			}
			result = append(result, lineResult...)
		}
	}
	
	return result
}

// GenerateArtWithColor converts input text to ASCII art with optional color support
func GenerateArtWithColor(text string, charMap map[rune][]string, substring, color string) string {
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

		// Generate ASCII art for this line with wrapping and color
		artLines := generateLineArtWithWrapAndColor(line, charMap, substring, color)
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

// generateLineArtWithWrapColorAndAlignment generates ASCII art for a line with terminal width wrapping, color, and alignment support
func generateLineArtWithWrapColorAndAlignment(text string, charMap map[rune][]string, substring, color, alignment string) []string {
	termWidth := getTerminalWidth()
	maxWidth := termWidth - 2
	
	if maxWidth < 10 {
		lines := generateSegmentWithColor(text, charMap, []struct{ start, end int }{}, 0, color)
		if alignment != "" {
			lines = applyAlignmentToLinesWithText(lines, alignment, termWidth, text)
		}
		return lines
	}

	// Find all substring occurrences
	var substringRanges []struct{ start, end int }
	if substring != "" {
		for i := 0; i <= len(text)-len(substring); i++ {
			if text[i:i+len(substring)] == substring {
				substringRanges = append(substringRanges, struct{ start, end int }{i, i + len(substring)})
			}
		}
	}

	// Calculate total width needed
	totalWidth := 0
	for _, char := range text {
		if charLines, exists := charMap[char]; exists && len(charLines) > 0 {
			totalWidth += len(charLines[0])
		}
	}

	// If text fits on one line, no need to wrap
	if totalWidth <= maxWidth {
		lines := generateSegmentWithColor(text, charMap, substringRanges, 0, color)
		if alignment != "" {
			lines = applyAlignmentToLinesWithText(lines, alignment, termWidth, text)
		}
		return lines
	}

	// Calculate optimal segments for even distribution
	numLines := (totalWidth + maxWidth - 1) / maxWidth // Ceiling division
	targetWidth := totalWidth / numLines

	// Generate segments with more even distribution
	var segments [][]string
	currentText := ""
	currentWidth := 0
	textOffset := 0
	
	for _, char := range text {
		charLines, exists := charMap[char]
		if !exists {
			continue
		}
		
		charWidth := 0
		if len(charLines) > 0 {
			charWidth = len(charLines[0])
		}
		
		// Use target width for more even distribution
		if currentWidth+charWidth > targetWidth && currentText != "" && len(segments) < numLines-1 {
			currentArt := generateSegmentWithColor(currentText, charMap, substringRanges, textOffset, color)
			segments = append(segments, currentArt)
			
			textOffset += len(currentText)
			currentText = string(char)
			currentWidth = charWidth
		} else {
			currentText += string(char)
			currentWidth += charWidth
		}
	}
	
	if currentText != "" {
		currentArt := generateSegmentWithColor(currentText, charMap, substringRanges, textOffset, color)
		segments = append(segments, currentArt)
	}

	// Apply alignment to all segments uniformly
	var allLines []string
	for _, segment := range segments {
		if alignment != "" {
			segment = applyAlignmentToLinesWithText(segment, alignment, termWidth, text)
		}
		allLines = append(allLines, segment...)
	}
	
	return allLines
}

// applyAlignmentToLinesWithText applies alignment to a set of lines with original text context
func applyAlignmentToLinesWithText(lines []string, alignment string, termWidth int, originalText string) []string {
	switch alignment {
	case "right":
		return alignRightConsistent(lines, termWidth)
	case "center":
		return alignCenterConsistent(lines, termWidth)
	case "justify":
		return alignJustifyConsistent(lines, termWidth, originalText)
	default:
		return lines
	}
}

// applyAlignmentToLines applies alignment to a set of lines
func applyAlignmentToLines(lines []string, alignment string, termWidth int) []string {
	return applyAlignmentToLinesWithText(lines, alignment, termWidth, "")
}

// generateLineArtWithWrapAndColor generates ASCII art for a line with terminal width wrapping and color support
func generateLineArtWithWrapAndColor(text string, charMap map[rune][]string, substring, color string) []string {
	termWidth := getTerminalWidth()
	// Reserve 2 characters for $ signs
	maxWidth := termWidth - 2
	
	if maxWidth < 10 {
		// Terminal too narrow, use original method
		return generateSegmentWithColor(text, charMap, []struct{ start, end int }{}, 0, color)
	}

	// Find all substring occurrences in the original text first
	var substringRanges []struct{ start, end int }
	if substring != "" {
		for i := 0; i <= len(text)-len(substring); i++ {
			if text[i:i+len(substring)] == substring {
				substringRanges = append(substringRanges, struct{ start, end int }{i, i + len(substring)})
			}
		}
	}

	var allLines []string
	currentText := ""
	currentWidth := 0
	textOffset := 0 // Track position in original text
	
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
			// Generate art for current text segment with color
			currentArt := generateSegmentWithColor(currentText, charMap, substringRanges, textOffset, color)
			allLines = append(allLines, currentArt...)
			
			// Update offset and start new line
			textOffset += len(currentText)
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
		currentArt := generateSegmentWithColor(currentText, charMap, substringRanges, textOffset, color)
		allLines = append(allLines, currentArt...)
	}
	
	return allLines
}

// generateSegmentWithColor generates ASCII art for a text segment with color support based on original text positions
func generateSegmentWithColor(segmentText string, charMap map[rune][]string, substringRanges []struct{ start, end int }, segmentOffset int, color string) []string {
	if segmentText == "" {
		return []string{""}
	}

	// Initialize 8 lines for the ASCII art
	artLines := make([]string, 8)

	// Get color code if color is specified
	var colorCode, resetCode string
	if color != "" {
		if code, exists := colorMap[strings.ToLower(color)]; exists {
			colorCode = code
			resetCode = colorMap["reset"]
		}
	}

	// Process each character in the segment
	for charPos, char := range segmentText {
		if charLines, exists := charMap[char]; exists {
			// Calculate position in original text
			originalPos := segmentOffset + charPos
			
			// Check if this character should be colored
			shouldColor := false
			if len(substringRanges) == 0 && colorCode != "" {
				// Color entire output
				shouldColor = true
			} else {
				// Check if character is within any substring range
				for _, r := range substringRanges {
					if originalPos >= r.start && originalPos < r.end {
						shouldColor = true
						break
					}
				}
			}

			// Add each line of the character to the corresponding art line
			for i := 0; i < 8; i++ {
				if i < len(charLines) {
					if shouldColor {
						// Check if we need to start or end color
						prevCharColored := charPos > 0 && isPositionInRanges(segmentOffset+charPos-1, substringRanges)
						nextCharColored := charPos < len(segmentText)-1 && isPositionInRanges(segmentOffset+charPos+1, substringRanges)
						
						if !prevCharColored {
							// Start of colored section
							artLines[i] += colorCode
						}
						artLines[i] += charLines[i]
						if !nextCharColored {
							// End of colored section
							artLines[i] += resetCode
						}
					} else {
						artLines[i] += charLines[i]
					}
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

// ApplyAlignment applies the specified alignment to ASCII art lines
func ApplyAlignment(artLines []string, alignment string) []string {
	if alignment == "left" || alignment == "" {
		// Left alignment is the default - no changes needed
		return artLines
	}

	// Get current terminal width for alignment calculations
	termWidth := getTerminalWidth()
	if termWidth < 10 {
		// Terminal too narrow for alignment
		return artLines
	}

	// Apply the specified alignment
	switch alignment {
	case "right":
		return alignRightConsistent(artLines, termWidth)
	case "center":
		return alignCenterConsistent(artLines, termWidth)
	case "justify":
		return alignJustifyConsistent(artLines, termWidth, "")
	default:
		return artLines
	}
}

// isPositionInRanges checks if a position is within any of the given ranges
func isPositionInRanges(pos int, ranges []struct{ start, end int }) bool {
	for _, r := range ranges {
		if pos >= r.start && pos < r.end {
			return true
		}
	}
	return false
}

// alignRightConsistent aligns all ASCII art lines consistently to the right
func alignRightConsistent(artLines []string, termWidth int) []string {
	var result []string
	
	for _, line := range artLines {
		if line == "" || line == "$" {
			result = append(result, line)
			continue
		}
		
		// Remove the trailing $ to get actual content
		content := strings.TrimSuffix(line, "$")
		
		// Calculate visual length (excluding ANSI color codes)
		visualLen := getVisualLength(content)
		
		// Calculate padding for right alignment
		padding := termWidth - visualLen - 1 // -1 for the $ at the end
		if padding < 0 {
			padding = 0
		}
		
		result = append(result, strings.Repeat(" ", padding)+content+"$")
	}
	
	return result
}

// getVisualLength returns the visual length of a string, excluding ANSI color codes
func getVisualLength(s string) int {
	visualLen := 0
	inEscape := false
	
	for i := 0; i < len(s); i++ {
		if s[i] == '\033' && i+1 < len(s) && s[i+1] == '[' {
			inEscape = true
			i++ // skip the '['
		} else if inEscape && s[i] == 'm' {
			inEscape = false
		} else if !inEscape {
			visualLen++
		}
	}
	
	return visualLen
}

// alignCenterConsistent centers all ASCII art lines consistently
func alignCenterConsistent(artLines []string, termWidth int) []string {
	var result []string
	
	for _, line := range artLines {
		if line == "" || line == "$" {
			result = append(result, line)
			continue
		}
		
		// Remove the trailing $ to get actual content
		content := strings.TrimSuffix(line, "$")
		
		// Calculate visual length (excluding ANSI color codes)
		visualLen := getVisualLength(content)
		
		// Calculate padding for center alignment
		padding := 0
		if visualLen < termWidth {
			totalPadding := termWidth - visualLen - 1 // -1 for the $ at the end
			padding = totalPadding / 2
		}
		
		result = append(result, strings.Repeat(" ", padding)+content+"$")
	}
	
	return result
}

// alignJustifyConsistent distributes all ASCII art lines consistently
// For single word: left-aligned. For multiple words: first word left, last word right
func alignJustifyConsistent(artLines []string, termWidth int, originalText string) []string {
	// Count words in original text
	words := strings.Fields(originalText)
	
	// If single word or empty, use left alignment
	if len(words) <= 1 {
		return artLines // Left-aligned (default)
	}
	
	// For multiple words, align to fill width (first word left, last word right)
	var result []string
	
	for _, line := range artLines {
		if line == "" || line == "$" {
			result = append(result, line)
			continue
		}
		
		// Remove the trailing $ to get actual content
		content := strings.TrimSuffix(line, "$")
		
		// Calculate visual length (excluding ANSI color codes)
		visualLen := getVisualLength(content)
		
		// Calculate padding for right alignment (to push to far right)
		padding := termWidth - visualLen - 1 // -1 for the $ at the end
		if padding < 0 {
			padding = 0
		}
		
		result = append(result, strings.Repeat(" ", padding)+content+"$")
	}
	
	return result
}