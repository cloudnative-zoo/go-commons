package gitlab

// GetUserID returns the user ID of the authenticated user.
func (s *Service) GetUserID() (int, error) {
	user, _, err := s.client.Users.CurrentUser()
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
