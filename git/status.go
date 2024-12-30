package git

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

type StatusResult struct {
	Added    []string
	Modified []string
	Deleted  []string
}

// Status represents the status of a git repository.
func (s *Service) Status() (*StatusResult, error) {
	// Get the working tree of the repository.
	worktree, err := s.repo.Worktree()
	if err != nil {
		return nil, fmt.Errorf("failed to access worktree for repository: %w", err)
	}

	// Get the status of the repository.
	status, err := worktree.Status()
	if err != nil {
		return nil, fmt.Errorf("failed to get status of repository: %w", err)
	}

	if status.IsClean() {
		// No changes to commit; silently return.
		return nil, nil
	}

	// Parse the status result.
	result := &StatusResult{}
	for file, flags := range status {
		switch {
		case flags.Staging == git.Unmodified: // Unmodified
			result.Modified = append(result.Modified, file)
		case flags.Staging == git.Added: // Added
			result.Added = append(result.Added, file)
		case flags.Staging == git.Deleted: // Deleted
			result.Deleted = append(result.Deleted, file)
		default:
			return nil, fmt.Errorf("unknown status flag for file %s: %v", file, flags)
		}
	}

	return result, nil
}
