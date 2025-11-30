package main

import (
	"os/exec"
	"strings"
	"testing"
)

// TestEdgeCases tests all edge cases that auditors typically check
func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectEmpty bool
		contains    string
	}{
		// Basic edge cases
		{"empty string", "", true, ""},
		{"single space", " ", false, "$"},
		{"only newline", "\\n", false, "$"},
		{"double newline", "\\n\\n", false, "$"},
		
		// Character range tests
		{"first printable char (space)", " ", false, "$"},
		{"exclamation mark", "!", false, "|"},
		{"tilde (last char)", "~", false, "/"},
		
		// Number tests
		{"single digit", "1", false, "/ |"},
		{"all digits", "0123456789", false, "___"},
		
		// Letter tests  
		{"single uppercase", "A", false, "/\\"},
		{"single lowercase", "a", false, "__ _"},
		{"full uppercase alphabet", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", false, "/\\"},
		{"full lowercase alphabet", "abcdefghijklmnopqrstuvwxyz", false, "__ _"},
		
		// Special character tests
		{"underscore (fat)", "_", false, "|_______|"},
		{"dash", "-", false, "______"},
		{"brackets", "[]", false, "|"},
		{"parentheses", "()", false, "/"},
		{"braces", "{}", false, "|"},
		
		// Multi-line tests
		{"hello newline world", "Hello\\nWorld", false, "_    _"},
		{"empty line between", "Hi\\n\\nBye", false, "$"},
		{"multiple newlines", "A\\n\\n\\nB", false, "/\\"},
		
		// Mixed content tests
		{"letters and numbers", "Hello123", false, "_    _"},
		{"letters and symbols", "Hello!", false, "_    _"},
		{"numbers and symbols", "123!", false, "/ |"},
		
		// Long strings
		{"very long string", "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", false, "/\\"},
		
		// Punctuation
		{"all punctuation", "!@#$%^&*()_+-=[]{}|;':\",./<>?", false, "|"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", "./cmd/ascii-art", tt.input)
			output, err := cmd.Output()
			
			if err != nil {
				t.Errorf("Command failed: %v", err)
				return
			}
			
			outputStr := string(output)
			
			if tt.expectEmpty {
				if outputStr != "" {
					t.Errorf("Expected empty output, got: %s", outputStr)
				}
			} else {
				if outputStr == "" {
					t.Error("Expected non-empty output")
				}
				
				if tt.contains != "" && !strings.Contains(outputStr, tt.contains) {
					t.Errorf("Output should contain '%s'", tt.contains)
				}
				
				// Check $ signs are present (except for empty input)
				if tt.input != "" && !strings.Contains(outputStr, "$") {
					t.Error("Output should contain $ signs at line ends")
				}
			}
		})
	}
}

// TestInvalidInputs tests how the program handles invalid inputs
func TestInvalidInputs(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"no arguments", []string{}},
		{"too many arguments", []string{"arg1", "arg2", "arg3"}},
		{"three arguments", []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", "./cmd/ascii-art")
			if len(tt.args) > 0 {
				cmd.Args = append(cmd.Args, tt.args...)
			}
			
			output, _ := cmd.Output()
			
			// Should produce no output for invalid number of arguments
			if len(output) > 0 {
				t.Errorf("Expected no output for invalid args, got: %s", string(output))
			}
		})
	}
}

// TestOutputFormat tests the exact output format requirements
func TestOutputFormat(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		checks []string
	}{
		{
			name:  "dollar signs at line ends",
			input: "A",
			checks: []string{"$"},
		},
		{
			name:  "8 lines for single character",
			input: "A", 
			checks: []string{"8_lines"},
		},
		{
			name:  "newline creates empty line with $",
			input: "A\\nB",
			checks: []string{"empty_line_dollar"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", "./cmd/ascii-art", tt.input)
			output, err := cmd.Output()
			
			if err != nil {
				t.Errorf("Command failed: %v", err)
				return
			}
			
			lines := strings.Split(string(output), "\n")
			
			for _, check := range tt.checks {
				switch check {
				case "8_lines":
					// Should have 8 lines for ASCII art (plus empty line at start)
					if len(lines) < 8 {
						t.Errorf("Expected at least 8 lines, got %d", len(lines))
					}
				case "empty_line_dollar":
					// Check for empty lines with $ in multi-line input
					if strings.Contains(tt.input, "\\n") {
						found := false
						for _, line := range lines {
							if line == "$" {
								found = true
								break
							}
						}
						if !found {
							t.Error("Expected to find empty line with $ in multi-line input")
						}
					}
				}
			}
		})
	}
}