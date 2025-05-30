package github

import (
	"errors"
	"fmt"
)

const (
	// DefaultPaginationMaxLimit is the default maximum number of items to request per page when paginating through results.
	DefaultPaginationMaxLimit = 50
)

// New initializes a new GitHub Service instance with the provided opts.
// At least one option must configure a GitHub client (e.g., WithToken).
func New(opts ...Options) (*Service, error) {
	s := &Service{}

	// Apply each option to configure the service.
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	// Ensure the GitHub client is configured.
	if s.client == nil {
		return nil, errors.New("missing GitHub client configuration")
	}

	// Set the default maximum number of items to request per page when paginating through results.
	if s.paginationMaxLimit == 0 {
		s.paginationMaxLimit = DefaultPaginationMaxLimit
	}

	return s, nil
}
