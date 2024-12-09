package util

import (
	"fmt"
	"os"
)

func ValidateAndCreateDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0o600); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	return nil
}
