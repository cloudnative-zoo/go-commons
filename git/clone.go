package git

import (
	"context"
	"fmt"

	"github.com/go-git/go-git/v5"
)

func (s *Service) Clone(ctx context.Context) error {
	cloneOptions := &git.CloneOptions{
		URL:               s.url,
		Auth:              s.auth,
		Progress:          s.progress,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	}

	_, err := git.PlainCloneContext(ctx, s.path, false, cloneOptions)
	if err != nil {
		return fmt.Errorf("failed to clone repository '%s' to '%s' error: %w", s.url, s.path, err)
	}

	return nil
}
