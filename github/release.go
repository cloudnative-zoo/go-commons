package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v69/github"
)

// GetLatestRelease fetches the latest release from a GitHub repository.
func (s *Service) GetLatestRelease(ctx context.Context, owner, repo string) (*github.RepositoryRelease, error) {
	release, _, err := s.client.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest release for %s/%s: %w", owner, repo, err)
	}
	return release, nil
}

// GetReleaseByTag fetches a specific release from a GitHub repository.
func (s *Service) GetReleaseByTag(ctx context.Context, owner, repo, tag string) (*github.RepositoryRelease, error) {
	release, _, err := s.client.Repositories.GetReleaseByTag(ctx, owner, repo, tag)
	if err != nil {
		return nil, fmt.Errorf("failed to get release %s for %s/%s: %w", tag, owner, repo, err)
	}
	return release, nil
}
