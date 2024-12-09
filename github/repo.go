package github

import (
	"context"

	"github.com/google/go-github/v67/github"
)

func (s *Service) ListOrganizationRepositories(ctx context.Context, org string) ([]*github.Repository, error) {
	return s.listRepositories(func(opts github.ListOptions) ([]*github.Repository, *github.Response, error) {
		orgOpts := &github.RepositoryListByOrgOptions{ListOptions: opts}
		return s.client.Repositories.ListByOrg(ctx, org, orgOpts)
	})
}

func (s *Service) ListUserRepositories(ctx context.Context) ([]*github.Repository, error) {
	return s.listRepositories(func(opts github.ListOptions) ([]*github.Repository, *github.Response, error) {
		userOpts := &github.RepositoryListByAuthenticatedUserOptions{
			Visibility:  "all",
			Affiliation: "owner",
		}
		return s.client.Repositories.ListByAuthenticatedUser(ctx, userOpts)
	})
}

func (s *Service) listRepositories(listFunc func(github.ListOptions) ([]*github.Repository, *github.Response, error)) ([]*github.Repository, error) {
	var repositories []*github.Repository
	opts := github.ListOptions{PerPage: 50}
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
