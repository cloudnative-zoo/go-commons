package git

import (
	"errors"
	"fmt"
)

// Status represents the status of a git repository.
func (s *Service) Status() error {
	// Get the working tree of the repository.
	worktree, err := s.repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to access worktree for repository: %w", err)
	}

	// Get the status of the repository.
	status, err := worktree.Status()
	if err != nil {
		return fmt.Errorf("failed to get status of repository: %w", err)
	}

	if status.IsClean() {
		// No changes to commit; silently return.
		return nil
	}

	return errors.New("repository has uncommitted changes")
}
