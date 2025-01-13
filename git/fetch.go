package git

import (
	"context"
	"errors"
	"fmt"

	gogit "github.com/go-git/go-git/v5"
)

// Fetch fetches updates from the remote repository.
// It supports forced fetches, shallow cloning, authentication, and progress reporting.
func (s *Service) Fetch(ctx context.Context) error {
	fetchOptions := &gogit.FetchOptions{
		Force:    true,       // Force overwriting changes
		Depth:    1,          // Perform a shallow fetch
		Auth:     s.auth,     // Authentication credentials
		Progress: s.progress, // Progress reporter (e.g., os.Stdout)
	}

	// Perform the fetch operation with context.
	err := s.repo.FetchContext(ctx, fetchOptions)
	if err != nil && !errors.Is(err, gogit.NoErrAlreadyUpToDate) {
		// Handle all errors except when the repository is already up to date.
		return fmt.Errorf("failed to fetch updates from repository: %w", err)
	}

	return nil
}
