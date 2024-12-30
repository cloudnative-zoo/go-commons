package git

// Open opens a repository at the specified path.
// If the repository does not exist, it returns an error.
// func (s *Service) Open() error {
// 	// Try to open the existing repository.
// 	var err error
// 	s.repo, err = git.PlainOpen(s.path)
// 	if err != nil {
// 		return fmt.Errorf("failed to open repository: %w", err)
// 	}
//
// 	return nil
// }
