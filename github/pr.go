package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v68/github"
)

// ListOpenPullRequests lists all pull requests for a repository.
// Returns a list of pull requests in the repository.
func (s *Service) ListOpenPullRequests(ctx context.Context, owner, repo string) ([]*github.PullRequest, error) {
	opts := &github.PullRequestListOptions{
		State: "open",
	}
	prs, _, err := s.client.PullRequests.List(ctx, owner, repo, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to list pull requests: %w", err)
	}
	return prs, nil
}

// GetPullRequest retrieves a pull request by its number.
// Returns the pull request with the specified number.
func (s *Service) GetPullRequest(ctx context.Context, owner, repo string, number int) (*github.PullRequest, error) {
	pr, _, err := s.client.PullRequests.Get(ctx, owner, repo, number)
	if err != nil {
		return nil, fmt.Errorf("failed to get pull request: %w", err)
	}
	return pr, nil
}

// MergePullRequest merges a pull request by its number.
// Returns the merged pull request.
func (s *Service) MergePullRequest(ctx context.Context, owner, repo string, number int) (*github.PullRequestMergeResult, error) {
	opts := &github.PullRequestOptions{
		MergeMethod: "squash",
	}
	pr, _, err := s.client.PullRequests.Merge(ctx, owner, repo, number, "", opts)
	if err != nil {
		return nil, fmt.Errorf("failed to merge pull request: %w", err)
	}
	return pr, nil
}