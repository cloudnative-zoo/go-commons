package github

import (
	"errors"
	"fmt"
)

// New initializes a new GitHub Service instance with the provided options.
// At least one option must configure a GitHub client (e.g., WithToken).
func New(options ...Option) (*Service, error) {
	s := &Service{}

	// Apply each option to configure the service.
	for _, opt := range options {
		if err := opt(s); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	// Ensure the GitHub client is configured.
	if s.client == nil {
		return nil, errors.New("missing GitHub client configuration")
	}

	return s, nil
}
