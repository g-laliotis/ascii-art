package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art/internal/ascii"
)

func main() {
	var colorFlag, substring, text, outputFile, banner, alignFlag string
	banner = "standard" // default banner
	hasColorFlag := false

	// Parse arguments - extract flags first
	args := os.Args[1:]
	for i := len(args) - 1; i >= 0; i-- {
		arg := args[i]
		// Parse --output=filename flag
		if strings.HasPrefix(arg, "--output=") {
			outputFile = strings.TrimPrefix(arg, "--output=")
			args = append(args[:i], args[i+1:]...)
		// Parse --color=color flag
		} else if strings.HasPrefix(arg, "--color=") {
			colorFlag = strings.TrimPrefix(arg, "--color=")
			hasColorFlag = true
			args = append(args[:i], args[i+1:]...)
		// Parse --align=type flag
		} else if strings.HasPrefix(arg, "--align=") {
			alignFlag = strings.TrimPrefix(arg, "--align=")
			// Validate alignment type
			if !isValidAlignment(alignFlag) {
				printUsage()
				return
			}
			args = append(args[:i], args[i+1:]...)
		}
	}

	switch len(args) {
	case 1:
		// "text" -> use default standard banner
		text = args[0]
	case 2:
		if hasColorFlag {
			// "--color=red text banner" -> color entire text with banner
			text = args[0]
			banner = args[1]
		} else if strings.Contains(args[0], "=") && !strings.HasPrefix(args[0], "--") {
			// Invalid flag format (e.g., "color=red")
			printUsage()
			return
		} else {
			// "text banner" -> use specified banner
			text = args[0]
			banner = args[1]
		}
	case 3:
		if hasColorFlag {
			// "--color=red substring text banner" -> color specific substring with banner
			substring = args[0]
			text = args[1]
			banner = args[2]
		} else {
			// 3 args without color flag is invalid
			printUsage()
			return
		}
	default:
		// Too many or no arguments
		if len(args) > 0 {
			printUsage()
			return
		}
		return
	}

	if text == "" {
		return
	}

	// Load the specified banner
	bannerFile := "assets/" + banner + ".txt"
	charMap, err := ascii.LoadBanner(bannerFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading banner: %v\n", err)
		os.Exit(1)
	}

	// Generate ASCII art with color and alignment support
	result := ascii.GenerateArtWithColorAndAlignment(text, charMap, substring, colorFlag, alignFlag)
	if result != "" {
		// Save to file or print to stdout
		if outputFile != "" {
			if err := ascii.SaveToFile(outputFile, result+"\n"); err != nil {
				fmt.Fprintf(os.Stderr, "Error saving to file: %v\n", err)
				os.Exit(1)
			}
		} else {
			fmt.Print(result + "\n")
		}
	}
}

// isValidAlignment checks if the alignment type is valid
func isValidAlignment(align string) bool {
	validAlignments := []string{"left", "right", "center", "justify"}
	for _, valid := range validAlignments {
		if align == valid {
			return true
		}
	}
	return false
}

func printUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println("\nExample: go run . --align=right something standard")
}