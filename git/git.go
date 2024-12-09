package git

import (
	"context"

	"github.com/go-git/go-git/v5"
)

func (s *Service) CloneOrPull(ctx context.Context) error {
	var err error
	s.repo, err = git.PlainOpen(s.path)
	if err != nil {
		return s.Clone(ctx)
	}
	return s.Pull(ctx)
}
