package git

import (
	"fmt"

	gogit "github.com/go-git/go-git/v5"
)

type StatusChanges struct {
	Added    []string
	Modified []string
	Deleted  []string
}

// Status represents the status of a git repository.
func (s *Service) Status() (*StatusChanges, error) {
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
		return &StatusChanges{}, nil
	}

	// Parse the status result.
	result := &StatusChanges{}
	for file, flags := range status {
		switch {
		case flags.Worktree == gogit.Modified || flags.Staging == gogit.Modified || flags.Worktree == gogit.Renamed || flags.Staging == gogit.Renamed || flags.Staging == gogit.UpdatedButUnmerged || flags.Worktree == gogit.UpdatedButUnmerged:
			result.Modified = append(result.Modified, file)
		case flags.Staging == gogit.Added || flags.Worktree == gogit.Added || flags.Worktree == gogit.Copied || flags.Staging == gogit.Copied:
			result.Added = append(result.Added, file)
		case flags.Staging == gogit.Deleted || flags.Worktree == gogit.Deleted:
			result.Deleted = append(result.Deleted, file)
		default:
			return nil, fmt.Errorf("unknown status flag for file %s: %v", file, flags)
		}
	}

	return result, nil
}
