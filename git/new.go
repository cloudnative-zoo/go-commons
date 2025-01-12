package git

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
)

// New initializes a new Git service instance with the provided options.
// It tries to open an existing repository at the specified path. If the repository does not exist, it clones it from the provided URL.
// It validates required fields such as authentication, repository URL, and local path.
func New(ctx context.Context, opts ...Options) (*Service, error) {
	// Create a new Service instance.
	service := &Service{}

	// Apply provided options to configure the service.
	for _, opt := range opts {
		if err := opt(service); err != nil {
			return nil, errors.New("failed to apply options: " + err.Error())
		}
	}

	// Validate required fields.
	if service.auth == nil {
		return nil, errors.New("authentication is required: provide a token, SSH key, or SSH key path")
	}
	if service.url == "" {
		remoteURL, err := getRemoteURL(service.path)
		if err != nil {
			return nil, fmt.Errorf("repository URL is required: %w", err)
		}
		slog.With("path", service.path).With("url", remoteURL).Info("using remote URL")
		service.url = remoteURL
	}
	if service.path == "" {
		return nil, errors.New("repository path is required")
	}

	// Try to open the existing repository.
	repo, err := open(service.path)
	if err != nil {
		// check if cloneIfNotExist is set to true in the options, if so, clone the repository. Otherwise, return an error.
		if !service.cloneIfNotExist {
			return nil, fmt.Errorf("repository not found at path %s: %w", service.path, err)
		}

		slog.With("path", service.path).Error("repository not found; attempting to clone")
		if cloneErr := clone(ctx, service.path, service.url, service.auth, service.progress); cloneErr != nil {
			return nil, fmt.Errorf("failed to clone repository from URL %s to path %s: %w", service.url, service.path, cloneErr)
		}

		// Reopen the repository after cloning.
		repo, err = open(service.path)
		if err != nil {
			return nil, errors.New("failed to open repository after cloning: " + err.Error())
		}
	}

	// Assign the opened repository to the service.
	service.repo = repo
	return service, nil
}
