package git

import (
	"errors"
)

func New(options ...Option) (*Service, error) {
	s := &Service{}
	for _, opt := range options {
		if err := opt(s); err != nil {
			return nil, err
		}
	}

	if s.auth == nil {
		return nil, errors.New("either token, ssh key, or ssh key path must be provided")
	}
	if s.url == "" {
		return nil, errors.New("clone url is required")
	}
	if s.path == "" {
		return nil, errors.New("repo path is required")
	}

	return s, nil
}
