package main

import (
	"fmt"
	"os"

	"ascii-art/internal/ascii"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	text := os.Args[1]
	if text == "" {
		return
	}

	// Load the standard banner
	charMap, err := ascii.LoadBanner("assets/standard.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading banner: %v\n", err)
		os.Exit(1)
	}

	// Generate and print ASCII art
	result := ascii.GenerateArt(text, charMap)
	if result != "" {
		fmt.Print(result + "\n")
	}
}