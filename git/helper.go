package git

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp/sideband"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

func open(path string) (*git.Repository, error) {
	// Try to open the existing repository. If it does not exist, return an error.
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open repository: %w", err)
	}
	return repo, nil
}

func clone(ctx context.Context, path, url string, auth transport.AuthMethod, progress sideband.Progress) error {
	// Configure the clone operation.
	cloneOptions := &git.CloneOptions{
		URL:               url,
		Auth:              auth,
		Progress:          progress,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	}

	// Perform the clone operation with context.
	_, err := git.PlainCloneContext(ctx, path, false, cloneOptions)
	if err != nil {
		return fmt.Errorf("failed to clone repository '%s' to '%s': %w", url, path, err)
	}

	return nil
}

func getRemoteURL(path string) (string, error) {
	repo, err := open(path)
	if err != nil {
		return "", err
	}

	remote, err := repo.Remote("origin")
	if err != nil {
		return "", fmt.Errorf("failed to get remote 'origin': %w", err)
	}

	remoteConfig := remote.Config()
	urls := remoteConfig.URLs
	if len(urls) == 0 {
		return "", errors.New("no remote URL found for 'origin'")
	}

	return urls[0], nil
}
