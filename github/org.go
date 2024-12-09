package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v67/github"
)

func (s *Service) ListOrganizations(ctx context.Context) ([]*github.Organization, error) {
	var orgs []*github.Organization
	opts := &github.ListOptions{PerPage: 50}

	for {
		organizations, resp, err := s.client.Organizations.List(ctx, "", opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list organizations with error: %w", err)
		}

		orgs = append(orgs, organizations...)

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	return orgs, nil
}
