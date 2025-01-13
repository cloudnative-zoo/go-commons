package git

import (
	"context"
	"errors"
	"fmt"

	gogit "github.com/go-git/go-git/v5"
)

// Pull updates the local repository by pulling changes from the remote repository.
// It handles authentication, progress reporting, and skips updates if the repository is already up-to-date.
func (s *Service) Pull(ctx context.Context) error {
	// Get the working tree of the repository.
	worktree, err := s.repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to access worktree for repository: %w", err)
	}

	// Configure pull opts.
	pullOptions := &gogit.PullOptions{
		RemoteName: "origin",   // Default remote name
		Auth:       s.auth,     // Authentication credentials
		Progress:   s.progress, // Progress reporting (e.g., os.Stdout)
	}

	// Perform the pull operation with context.
	err = worktree.PullContext(ctx, pullOptions)
	switch {
	case errors.Is(err, gogit.NoErrAlreadyUpToDate):
		// No changes to pull; silently return.
		return nil
	case err != nil:
		// Return any other errors encountered during pull.
		return fmt.Errorf("failed to pull repository: %w", err)
	}

	return nil
}
