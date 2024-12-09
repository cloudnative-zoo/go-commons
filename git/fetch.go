package git

import (
	"context"
	"fmt"

	"github.com/go-git/go-git/v5"
)

func (s *Service) Fetch(ctx context.Context) error {
	fetchOptions := &git.FetchOptions{
		Force:    true,
		Depth:    1,
		Auth:     s.auth,
		Progress: s.progress,
	}

	err := s.repo.FetchContext(ctx, fetchOptions)
	if err != nil {
		return fmt.Errorf("failed to fetch from repository with %s", err.Error())
	}
	return nil
}
