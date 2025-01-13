package git

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp/sideband"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

func open(path string) (*gogit.Repository, error) {
	// Try to open the existing repository. If it does not exist, return an error.
	repo, err := gogit.PlainOpen(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open repository: %w", err)
	}
	return repo, nil
}

func clone(ctx context.Context, path, url string, auth transport.AuthMethod, progress sideband.Progress) error {
	// Configure the clone operation.
	cloneOptions := &gogit.CloneOptions{
		URL:               url,
		Auth:              auth,
		Progress:          progress,
		RecurseSubmodules: gogit.DefaultSubmoduleRecursionDepth,
	}

	// Perform the clone operation with context.
	_, err := gogit.PlainCloneContext(ctx, path, false, cloneOptions)
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

// validateCredentials performs a minimal fetch to verify that the credentials are valid.
// It uses an in-memory repository, adds a remote, and fetches from it with the provided Auth.
// If the fetch fails due to invalid credentials, this returns an error.
func validateCredentials(url string, auth transport.AuthMethod) error {
	// 1. Create a temporary, in-memory repository.
	//    (No working directory, so we pass `nil` as the second argument.)
	r, err := gogit.Init(memory.NewStorage(), nil)
	if err != nil {
		return fmt.Errorf("failed to init in-memory repo: %w", err)
	}

	// 2. Add a remote named "testAuth" pointing to the desired URL.
	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: "testAuth",
		URLs: []string{url},
	})
	if err != nil {
		return fmt.Errorf("failed to create remote for validation: %w", err)
	}

	// 3. Perform a fetch on the "testAuth" remote using the provided Auth.
	//    If credentials are invalid, we'll get an error back.
	err = r.Fetch(&gogit.FetchOptions{
		RemoteName: "testAuth",
		Auth:       auth,
	})
	if err != nil {
		// If the remote was already fetched or if the repo is empty, go-git might return
		// `gogit.NoErrAlreadyUpToDate`. That’s not necessarily a credentials' error.
		// So we can do a quick check for that:
		if errors.Is(err, gogit.NoErrAlreadyUpToDate) {
			return nil
		}
		return fmt.Errorf("fetch failed, possibly invalid credentials: %w", err)
	}

	// If we got here, credentials worked successfully.
	return nil
}
