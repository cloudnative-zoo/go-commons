package git

import (
	"fmt"
)

func (s *Service) Add(files []string) error {
	w, err := s.repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	for _, file := range files {
		_, err := w.Add(file)
		if err != nil {
			return fmt.Errorf("failed to add file '%s': %w", file, err)
		}
	}

	return nil
}
