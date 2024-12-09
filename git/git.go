package git

import (
	"context"

	"github.com/go-git/go-git/v5"
)

// CloneOrPull ensures the local repository is up-to-date.
// If the repository does not exist at the specified path, it performs a clone operation.
// If the repository exists, it performs a pull to fetch the latest changes.
func (s *Service) CloneOrPull(ctx context.Context) error {
	// Try to open the existing repository.
	var err error
	s.repo, err = git.PlainOpen(s.path)
	if err != nil {
		// Repository does not exist; perform a clone.
		return s.Clone(ctx)
	}

	// Repository exists; perform a pull.
	return s.Pull(ctx)
}
