package ascii

import (
	"os"
	"testing"
)

func TestSaveToFile(t *testing.T) {
	testContent := "test content\nline 2"
	testFile := "test_output.txt"

	// Clean up after test
	defer os.Remove(testFile)

	// Test saving to file
	err := SaveToFile(testFile, testContent)
	if err != nil {
		t.Errorf("SaveToFile() error = %v", err)
		return
	}

	// Verify file was created and content is correct
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Errorf("Failed to read test file: %v", err)
		return
	}

	if string(content) != testContent {
		t.Errorf("SaveToFile() content = %q, want %q", string(content), testContent)
	}
}
