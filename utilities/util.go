package utilities

import "strings"

func ContainsAny(key string, keys []string) bool {
	if key == "" {
		return false
	}
	for _, k := range keys {
		if key == k {
			return true
		}
	}
	return false
}

func MatchesAny(name string, patterns []string) bool {
	if name == "" {
		return false
	}
	for _, pattern := range patterns {
		if strings.Contains(name, pattern) {
			return true
		}
	}
	return false
}
