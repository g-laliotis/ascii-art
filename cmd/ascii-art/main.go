package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art/internal/ascii"
)

func main() {
	var colorFlag, substring, text string

	// Parse arguments
	switch len(os.Args) {
	case 2:
		text = os.Args[1]
	case 3:
		// --color=<color> "text"
		if strings.HasPrefix(os.Args[1], "--color=") {
			colorFlag = strings.TrimPrefix(os.Args[1], "--color=")
			text = os.Args[2]
		} else {
			printUsage()
			return
		}
	case 4:
		// --color=<color> <substring> "text"
		if strings.HasPrefix(os.Args[1], "--color=") {
			colorFlag = strings.TrimPrefix(os.Args[1], "--color=")
			substring = os.Args[2]
			text = os.Args[3]
		} else {
			printUsage()
			return
		}
	default:
		if len(os.Args) > 2 {
			printUsage()
		}
		return
	}

	if text == "" {
		return
	}

	// Load the standard banner
	charMap, err := ascii.LoadBanner("assets/standard.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading banner: %v\n", err)
		os.Exit(1)
	}

	// Generate ASCII art
	result := ascii.GenerateArt(text, charMap)
	if result != "" {
		// Apply color if flag is present
		if colorFlag != "" {
			lines := strings.Split(result, "\n")
			lines = ascii.ApplyColor(lines, substring, colorFlag, text, charMap)
			result = strings.Join(lines, "\n")
		}
		fmt.Print(result + "\n")
	}
}

func printUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println("\nEX: go run . --color=<color> <substring to be colored> \"something\"")
}