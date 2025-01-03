package util

import (
	"os"
	"path/filepath"
)

// CreateDir create new dir and sub dir if not exist.
func CreateDir(dirPath ...string) error {
	for _, dir := range dirPath {
		newDir := filepath.Join(".", dir)
		err := os.MkdirAll(newDir, 0o600) //nolint:mnd
		if err != nil {
			return err
		}
	}
	return nil
}

// RemoveDir is removing dir.
func RemoveDir(dirs ...string) error {
	for _, dir := range dirs {
		if dir != "" {
			err := os.RemoveAll(dir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// CleanDir is removing all files in dir.
func CleanDir(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		err := os.Remove(filepath.Join(dir, file.Name()))
		if err != nil {
			return err
		}
	}
	return nil
}
