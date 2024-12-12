// Package gitlab provides a Gitlab Service client for interacting with the Gitlab API.
package gitlab

import (
	"errors"
	"fmt"
)

const (
	// DefaultPaginationMaxLimit is the default maximum number of items to request per page when paginating through results.
	DefaultPaginationMaxLimit = 500
	// DefaultPage is the default page number to request when paginating through results.
	DefaultPage = 1
	// Sort by ascending order.
	Sort = "asc"
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

	// Set the default maximum number of items to request per page when paginating through results.
	if s.paginationMaxLimit == 0 {
		s.paginationMaxLimit = DefaultPaginationMaxLimit
	}

	return s, nil
}
