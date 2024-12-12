package gitlab

import (
	"fmt"

	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// ListUserGroups returns a list of groups for the authenticated user.
func (s *Service) ListUserGroups() ([]*gitlab.Group, error) {
	owned := true
	groups, _, err := s.client.Groups.ListGroups(&gitlab.ListGroupsOptions{
		ListOptions: *s.listOptions,

		Owned: &owned,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list groups: %w", err)
	}
	return groups, nil
}
