package git

import (
	"errors"
	"log/slog"

	"github.com/go-git/go-git/v5"
)

// New initializes a new Git service instance with the provided opts.
// It validates required fields like authentication, repository URL, and local path.
func New(opts ...Options) (*Service, error) {
	s := &Service{}

	// Apply each option to configure the service.
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}

	// Validate required fields.
	if s.auth == nil {
		return nil, errors.New("authentication is required: provide a token, SSH key, or SSH key path")
	}
	if s.url == "" {
		return nil, errors.New("clone URL is required")
	}
	if s.path == "" {
		return nil, errors.New("repository path is required")
	}

	// Try to open the existing repository.
	var err error
	s.repo, err = git.PlainOpen(s.path)
	if err != nil {
		slog.With("path", s.path).Debug("repository does not exist; consider git clone")
	}

	return s, nil
}
