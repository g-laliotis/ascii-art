package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	text := os.Args[1]
	if text == "" {
		return
	}

	// TODO: Implement ASCII art generation
	fmt.Println("ASCII-Art placeholder for:", text)
}