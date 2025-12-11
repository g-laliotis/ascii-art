package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art/internal/ascii"
)

func main() {
	var colorFlag, substring, text, outputFile string

	// Parse arguments
	args := os.Args[1:]
	for i, arg := range args {
		if strings.HasPrefix(arg, "--output=") {
			outputFile = strings.TrimPrefix(arg, "--output=")
			args = append(args[:i], args[i+1:]...)
			break
		}
	}

	switch len(args) {
	case 1:
		text = args[0]
	case 2:
		if strings.HasPrefix(args[0], "--color=") {
			colorFlag = strings.TrimPrefix(args[0], "--color=")
			text = args[1]
		} else {
			printUsage()
			return
		}
	case 3:
		if strings.HasPrefix(args[0], "--color=") {
			colorFlag = strings.TrimPrefix(args[0], "--color=")
			substring = args[1]
			text = args[2]
		} else {
			printUsage()
			return
		}
	default:
		if len(args) > 0 {
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

func printUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println("\nEX: go run . --color=<color> <substring to be colored> \"something\"")
}