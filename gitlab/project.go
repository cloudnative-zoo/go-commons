package gitlab

import (
	"fmt"

	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// ListOwnedProjects returns a list of projects for the authenticated user.
func (s *Service) ListOwnedProjects() ([]*gitlab.Project, error) {
	// Fetch the list of projects from the Gitlab API.
	owned := true
	projects, _, err := s.client.Projects.ListProjects(&gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: s.paginationMaxLimit,
			Page:    1,
		},

		Owned:      &owned,
		Membership: &owned,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list projects: %w", err)
	}
	return projects, nil
}
