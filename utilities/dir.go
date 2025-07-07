package utilities

import (
	"os"
	"path/filepath"
)

const (
	DirPermOwnerOnly os.FileMode = 0o700 // Full access for owner, no access for others
	DirPermAllRead   os.FileMode = 0o755 // Full access for owner, read/execute for others
	DirPermAllWrite  os.FileMode = 0o777 // Full access for everyone
)

// CreateDir creates new directories and subdirectories if they don't exist.
// If perm is nil, it uses DirPermAllRead as the default.
func CreateDir(perm *os.FileMode, dirPath ...string) error {
	if perm == nil {
		// Default permission if perm is nil
		defaultPerm := DirPermAllRead
		perm = &defaultPerm
	}

	for _, dir := range dirPath {
		newDir := filepath.Join(".", dir)
		err := os.MkdirAll(newDir, *perm)
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
