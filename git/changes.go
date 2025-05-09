package git

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/gobwas/glob"
)

// HasRepositoryChanges checks if the repository has any changes (staged or unstaged).
func (s *Service) HasRepositoryChanges() (bool, error) {
	worktree, err := s.repo.Worktree()
	if err != nil {
		return false, fmt.Errorf("unable to access worktree: %w", err)
	}

	status, err := worktree.Status()
	if err != nil {
		return false, fmt.Errorf("unable to retrieve repository status: %w", err)
	}

	for _, fileStatus := range status {
		if fileStatus.Worktree != git.Unmodified || fileStatus.Staging != git.Unmodified {
			return true, nil
		}
	}
	return false, nil
}

// HasStagedChanges checks if there are any staged changes in the repository.
func (s *Service) HasStagedChanges() (bool, error) {
	worktree, err := s.repo.Worktree()
	if err != nil {
		return false, fmt.Errorf("unable to access worktree: %w", err)
	}

	status, err := worktree.Status()
	if err != nil {
		return false, fmt.Errorf("unable to retrieve repository status: %w", err)
	}

	for _, fileStatus := range status {
		if fileStatus.Staging != git.Unmodified {
			return true, nil
		}
	}
	return false, nil
}

// StageChanges stages all changes in the repository, excluding files matching the provided patterns.
func (s *Service) StageChanges(excludePatterns string) error {
	worktree, err := s.repo.Worktree()
	if err != nil {
		return fmt.Errorf("unable to access worktree: %w", err)
	}

	// Stage all changes.
	if err := worktree.AddGlob("."); err != nil {
		return fmt.Errorf("failed to stage changes: %w", err)
	}

	// Handle exclusion patterns if provided.
	if excludePatterns != "" {
		patternList := strings.Split(excludePatterns, ",")
		for _, pattern := range patternList {
			pattern = strings.TrimSpace(pattern)
			if pattern == "" {
				continue
			}

			matcher, err := glob.Compile(pattern)
			if err != nil {
				return fmt.Errorf("invalid exclude pattern '%s': %w", pattern, err)
			}

			status, err := worktree.Status()
			if err != nil {
				return fmt.Errorf("unable to retrieve repository status: %w", err)
			}

			for filePath := range status {
				if matcher.Match(filePath) {
					if err := worktree.RemoveGlob(filePath); err != nil {
						return fmt.Errorf("failed to unstage file '%s': %w", filePath, err)
					}
				}
			}
		}
	}

	return nil
}
