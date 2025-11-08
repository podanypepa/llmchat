// Package imagetools provides utilities for handling image data.
package imagetools

import (
	"encoding/base64"
	"fmt"
	"os"
)

// SaveImage decodes a base64 string and saves it as an image file.
func SaveImage(base64Data, filename string) error {
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return fmt.Errorf("error decoding base64 data: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}
