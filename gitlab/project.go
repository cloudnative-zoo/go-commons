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
		ListOptions: *s.listOptions,

		Owned:      &owned,
		Membership: &owned,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list projects: %w", err)
	}
	return projects, nil
}

// ListUserProjects returns a list of projects for the authenticated user.
func (s *Service) ListUserProjects() ([]*gitlab.Project, error) {
	userID, err := s.GetUserID()
	if err != nil {
		return nil, err
	}

	projects, _, err := s.client.Projects.ListUserProjects(userID, &gitlab.ListProjectsOptions{
		ListOptions: *s.listOptions,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list projects: %w", err)
	}
	return projects, nil
}
