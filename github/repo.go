package github

import (
	"context"
	"errors"

	"github.com/google/go-github/v70/github"
)

// ListRepositories retrieves repositories for the given target (user or organization).
// If the target is an organization, it fetches organization repositories.
// If the target is empty, it fetches all repositories for the authenticated user.
func (s *Service) ListRepositories(ctx context.Context, target string) ([]*github.Repository, error) {
	if target == "" {
		return s.ListUserRepositories(ctx)
	}

	isOrg, err := s.isOrganization(ctx, target)
	if err != nil {
		return nil, err
	}

	if isOrg {
		return s.ListOrganizationRepositories(ctx, target)
	}

	return nil, errors.New("specify a valid organization or leave target empty for user repositories")
}

// ListOrganizationRepositories fetches all repositories for a given organization.
func (s *Service) ListOrganizationRepositories(ctx context.Context, org string) ([]*github.Repository, error) {
	return s.listRepositories(func(opts github.ListOptions) ([]*github.Repository, *github.Response, error) {
		orgOpts := &github.RepositoryListByOrgOptions{ListOptions: opts}
		return s.client.Repositories.ListByOrg(ctx, org, orgOpts)
	})
}

// ListUserRepositories fetches all repositories owned by the authenticated user.
func (s *Service) ListUserRepositories(ctx context.Context) ([]*github.Repository, error) {
	return s.listRepositories(func(opts github.ListOptions) ([]*github.Repository, *github.Response, error) {
		userOpts := &github.RepositoryListByAuthenticatedUserOptions{
			Visibility:  "all",
			Affiliation: "owner",
		}
		return s.client.Repositories.ListByAuthenticatedUser(ctx, userOpts)
	})
}

// listRepositories is a helper function for paginated repository fetching.
func (s *Service) listRepositories(listFunc func(github.ListOptions) ([]*github.Repository, *github.Response, error)) ([]*github.Repository, error) {
	var repositories []*github.Repository
	opts := github.ListOptions{PerPage: s.paginationMaxLimit}

	for {
		repos, resp, err := listFunc(opts)
		if err != nil {
			return nil, err
		}
		repositories = append(repositories, repos...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return repositories, nil
}
