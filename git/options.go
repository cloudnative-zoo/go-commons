package git

import (
	"errors"

	"github.com/go-git/go-git/v5/plumbing/protocol/packp/sideband"

	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

type Option func(s *Service) error

func WithToken(token string) Option {
	return func(s *Service) error {
		if token == "" {
			return errors.New("token cannot be empty")
		}
		s.auth = &http.BasicAuth{
			Username: "git",
			Password: token,
		}
		return nil
	}
}

func WithURL(url string) Option {
	return func(s *Service) error {
		if url == "" {
			return errors.New("url cannot be empty")
		}
		s.url = url
		return nil
	}
}

func WithSSHKeyPath(path string, passphrase string) Option {
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

func WithSSHKey(key []byte, passphrase string) Option {
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

func WithRepoPath(path string) Option {
	return func(s *Service) error {
		if path == "" {
			return errors.New("repo path cannot be empty")
		}
		s.path = path
		return nil
	}
}

func WithBranch(branch string) Option {
	return func(s *Service) error {
		s.branch = branch
		return nil
	}
}

func WithProgress(progress sideband.Progress) Option {
	return func(s *Service) error {
		s.progress = progress
		return nil
	}
}
