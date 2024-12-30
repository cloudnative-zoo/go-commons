package git

import (
	"context"
	"fmt"
)

// Diff returns the changes between the working directory and the index.
// It returns the changes as a string in the format of `git diff --cached --name-status`.
func (s *Service) Diff(ctx context.Context) (string, error) {
	// Get the working tree of the repository.
	_, err := s.repo.Worktree()
	if err != nil {
		return "", fmt.Errorf("failed to access worktree for repository: %w", err)
	}

	// TODO: Implement the diff operation.
	return "", nil
}
