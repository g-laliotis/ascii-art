package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantErr  bool
		contains string
	}{
		{
			name:     "empty string",
			args:     []string{""},
			wantErr:  false,
			contains: "",
		},
		{
			name:     "single character",
			args:     []string{"A"},
			wantErr:  false,
			contains: "/\\",
		},
		{
			name:     "simple word",
			args:     []string{"Hello"},
			wantErr:  false,
			contains: "_    _",
		},
		{
			name:     "with newline",
			args:     []string{"Hello\\nWorld"},
			wantErr:  false,
			contains: "$",
		},
		{
			name:     "numbers",
			args:     []string{"123"},
			wantErr:  false,
			contains: "/ |",
		},
		{
			name:     "special characters",
			args:     []string{"!@#"},
			wantErr:  false,
			contains: "|",
		},
		{
			name:     "underscore fat",
			args:     []string{"_"},
			wantErr:  false,
			contains: "|_______|",
		},
		{
			name:     "mixed case",
			args:     []string{"HeLLo"},
			wantErr:  false,
			contains: "_    _",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", "./cmd/ascii-art", tt.args[0])
			output, err := cmd.Output()
			
			if tt.wantErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			
			if tt.contains != "" && !strings.Contains(string(output), tt.contains) {
				t.Errorf("Output should contain '%s', got: %s", tt.contains, string(output))
			}
		})
	}
}

func TestMainEdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectUsage bool
	}{
		{"no arguments", []string{}, false},
		{"invalid flag", []string{"color=red", "test"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", "./cmd/ascii-art")
			cmd.Args = append(cmd.Args, tt.args...)
			output, _ := cmd.Output()
			
			if tt.expectUsage {
				if !strings.Contains(string(output), "Usage:") {
					t.Errorf("Expected usage message for %s, got: %s", tt.name, string(output))
				}
			} else {
				if len(output) > 0 {
					t.Errorf("Expected no output for %s, got: %s", tt.name, string(output))
				}
			}
		})
	}
}