package git

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
)

func (s *Service) Pull(ctx context.Context) error {
	worktree, err := s.repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree for repository with %s", err.Error())
	}

	pullOptions := &git.PullOptions{
		RemoteName: "origin",
		Auth:       s.auth,
		Progress:   s.progress,
	}

	err = worktree.PullContext(ctx, pullOptions)
	switch {
	case errors.Is(err, git.NoErrAlreadyUpToDate):
		return nil
	case err != nil:
		return fmt.Errorf("failed to pull repository: %w", err)
	}
	return nil
}
