package git

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

var ErrNoStagedChanges = errors.New("no staged changes to commit")

// Commit creates a new commit with the given message.
// It auto-stages all changes when none are already staged.
func (s *Service) Commit(msg, commiterName, commiterEmail string) error {
	wt, err := s.repo.Worktree()
	if err != nil {
		return fmt.Errorf("access worktree: %w", err)
	}

	staged, err := hasStaged(wt)
	if err != nil {
		return fmt.Errorf("check staged changes: %w", err)
	}

	if !staged {
		if err := wt.AddGlob("."); err != nil {
			return fmt.Errorf("stage all changes: %w", err)
		}
		staged, err = hasStaged(wt)
		if err != nil {
			return fmt.Errorf("re-check staged changes: %w", err)
		}
		if !staged {
			return ErrNoStagedChanges
		}
	}

	_, err = wt.Commit(msg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  commiterName,
			Email: commiterEmail,
			When:  time.Now(),
		},
	})
	if err != nil {
		return fmt.Errorf("create commit: %w", err)
	}
	return nil
}

// hasStaged returns true if there is at least one staged change.
func hasStaged(wt *git.Worktree) (bool, error) {
	status, err := wt.Status()
	if err != nil {
		return false, err
	}
	for _, st := range status {
		if st.Staging != git.Unmodified {
			return true, nil
		}
	}
	return false, nil
}
