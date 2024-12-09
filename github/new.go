package github

import (
	"fmt"
)

func New(options ...Option) (*Service, error) {
	s := &Service{}
	for _, opt := range options {
		if err := opt(s); err != nil {
			return nil, err
		}
	}

	if s.client == nil {
		return nil, fmt.Errorf("github client is required")
	}

	return s, nil
}
