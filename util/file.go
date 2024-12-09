package util

import (
	"os"
	"path/filepath"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func DefaultSSHKey() string {
	defaultIDRSA := filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa")
	defaultIDED25519 := filepath.Join(os.Getenv("HOME"), ".ssh", "id_ed25519")
	if FileExists(defaultIDRSA) {
		return defaultIDRSA
	} else if FileExists(defaultIDED25519) {
		return defaultIDED25519
	} else {
		return ""
	}
}
