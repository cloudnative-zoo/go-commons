package git

import (
	"fmt"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func (s *Service) CheckoutBranch(name string, createIfNotExists bool) error {
	w, err := s.repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	branch := plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", name))
	err = w.Checkout(&gogit.CheckoutOptions{
		Branch: branch,
		Create: createIfNotExists,
	})
	if err != nil {
		return fmt.Errorf("failed to checkout branch '%s': %w", name, err)
	}

	return nil
}

func (s *Service) ListLocalBranches() ([]string, error) {
	branches, err := s.repo.Branches()
	if err != nil {
		return nil, fmt.Errorf("failed to get branches: %w", err)
	}

	var names []string
	err = branches.ForEach(func(ref *plumbing.Reference) error {
		names = append(names, ref.Name().Short())
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list branches: %w", err)
	}

	return names, nil
}

func (s *Service) ListRemoteBranches() ([]string, error) {
	remote, err := s.repo.Remote("origin")
	if err != nil {
		return nil, fmt.Errorf("failed to get remote 'origin': %w", err)
	}

	// Make sure to include s.auth (whatever you store in your Service struct) so go-git can authenticate
	refs, err := remote.List(&gogit.ListOptions{
		Auth: s.auth, // <-- Pass your SSH or HTTP Auth object here
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list remote branches: %w", err)
	}

	var names []string
	for _, ref := range refs {
		names = append(names, ref.Name().Short())
	}

	return names, nil
}
