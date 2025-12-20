package ascii

import (
	"os"
)

// SaveToFile saves the ASCII art output to a file
func SaveToFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}