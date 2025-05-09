package git

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/go-git/go-git/v5/plumbing/protocol/packp/sideband"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

// Options defines a functional option for configuring the Service.
type Options func(s *Service) error

// WithToken sets HTTP Basic Authentication using the provided personal access token.
func WithToken(token string) Options {
	return func(s *Service) error {
		if token == "" {
			return errors.New("token cannot be empty")
		}
		s.auth = &http.BasicAuth{
			Username: "faker", // Uses "faker" as the username for token authentication
			Password: token,
		}
		return nil
	}
}

// WithURL sets the repository clone URL.
func WithURL(url string) Options {
	return func(s *Service) error {
		if url == "" {
			return errors.New("url cannot be empty")
		}
		s.url = url
		return nil
	}
}

// WithSSHKeyPath sets SSH authentication using a private key file.
func WithSSHKeyPath(path string, passphrase string) Options {
	return func(s *Service) error {
		if path == "" {
			return errors.New("ssh key path cannot be empty")
		}
		sshAuth, err := ssh.NewPublicKeysFromFile("git", path, passphrase)
		if err != nil {
			return err
		}
		s.auth = sshAuth
		return nil
	}
}

// WithSSHKey sets SSH authentication using an in-memory private key.
func WithSSHKey(key []byte, passphrase string) Options {
	return func(s *Service) error {
		if len(key) == 0 {
			return errors.New("ssh key cannot be empty")
		}
		sshAuth, err := ssh.NewPublicKeys("git", key, passphrase)
		if err != nil {
			return err
		}
		s.auth = sshAuth
		return nil
	}
}

// WithRepoPath sets the local path where the repository will be cloned or updated.
func WithRepoPath(path string) Options {
	return func(s *Service) error {
		if path == "" {
			return errors.New("repo path cannot be empty")
		}
		// Ensure the path is absolute
		absPath, err := filepath.Abs(path)
		if err != nil {
			return fmt.Errorf("failed to get absolute path: %w", err)
		}
		s.path = absPath
		return nil
	}
}

// WithBranch sets the branch to be used for operations.
func WithBranch(branch string) Options {
	return func(s *Service) error {
		s.branch = branch
		return nil
	}
}

// WithProgress sets a progress writer for operations like clone and fetch.
func WithProgress(progress sideband.Progress) Options {
	return func(s *Service) error {
		s.progress = progress
		return nil
	}
}

// WithCloneIfNotExist sets whether to clone the repository if it does not exist locally.
func WithCloneIfNotExist(cloneIfNotExist bool) Options {
	return func(s *Service) error {
		s.cloneIfNotExist = cloneIfNotExist
		return nil
	}
}
