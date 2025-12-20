package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art/internal/ascii"
)

func main() {
	var colorFlag, substring, text, outputFile, banner string
	banner = "standard" // default banner

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
		// "text" -> use default standard banner
		text = args[0]
	case 2:
		if strings.HasPrefix(args[0], "--color=") {
			// "--color=red text" -> color entire text
			colorFlag = strings.TrimPrefix(args[0], "--color=")
			text = args[1]
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
		if strings.HasPrefix(args[0], "--color=") {
			// "--color=red substring text" -> color specific substring
			colorFlag = strings.TrimPrefix(args[0], "--color=")
			substring = args[1]
			text = args[2]
		} else {
			// 3 args without color flag is invalid
			printUsage()
			return
		}
	case 4:
		if strings.HasPrefix(args[0], "--color=") {
			// "--color=red substring text banner" -> color substring with specific banner
			colorFlag = strings.TrimPrefix(args[0], "--color=")
			substring = args[1]
			text = args[2]
			banner = args[3]
		} else {
			printUsage()
			return
		}
	default:
		// Too many or no arguments
		if len(args) > 0 {
			printUsage()
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

	// Generate ASCII art with color support
	result := ascii.GenerateArtWithColor(text, charMap, substring, colorFlag)
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

func printUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println("\nEX: go run . --color=<color> <substring to be colored> \"something\"")
}