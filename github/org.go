package github

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v72/github"
)

// ListOrganizations fetches all organizations associated with the authenticated user.
// Results are paginated and combined into a single list.
func (s *Service) ListOrganizations(ctx context.Context) ([]*github.Organization, error) {
	var orgs []*github.Organization
	opts := &github.ListOptions{PerPage: s.paginationMaxLimit}

	for {
		organizations, resp, err := s.client.Organizations.List(ctx, "", opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list organizations: %w", err)
		}

		orgs = append(orgs, organizations...)

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	return orgs, nil
}

// isOrganization determines if the specified target is a GitHub organization.
// It returns true if the target is an organization, false otherwise.
func (s *Service) isOrganization(ctx context.Context, target string) (bool, error) {
	user, _, err := s.client.Users.Get(ctx, target)
	if err != nil {
		// If the target does not exist, assume it's not an organization.
		if strings.Contains(err.Error(), "404") {
			return false, nil
		}
		return false, err
	}
	return user.GetType() == "Organization", nil
}
