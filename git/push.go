package git

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

// Push pushes commits to the remote repository.
func (s *Service) Push() error {
	// Get the current branch
	head, err := s.repo.Head()
	if err != nil {
		return fmt.Errorf("failed to get HEAD: %w", err)
	}

	// Create proper RefSpec
	refSpec := config.RefSpec(head.Name().String() + ":" + head.Name().String())

	// Push to remote
	err = s.repo.Push(&git.PushOptions{
		RefSpecs: []config.RefSpec{refSpec},
	})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return fmt.Errorf("failed to push: %w", err)
	}

	return nil
}
