package gitlab

import (
	"fmt"
	"log/slog"
	"os"

	gitlab "gitlab.com/gitlab-org/api/client-go"

	"github.com/cloudnative-zoo/go-commons/util"
)

// Options defines a function signature for configuring a Gitlab Service instance.
type Options func(*Service) error

// WithToken sets up authentication for the Gitlab client using a personal access token.
// The token can be provided directly or sourced from environment variables.
func WithToken(token string) Options {
	return func(s *Service) error {
		if token == "" {
			// Fetch token from environment variables if not provided.
			token = util.GetEnv("GITLAB_TOKEN", "GITLAB_API_TOKEN")
			if token == "" {
				slog.Error("A valid token must be provided directly or via environment variables (GITLAB_TOKEN, GITLAB_API_TOKEN).")
				os.Exit(1)
			}
		}

		var err error
		// Initialize the Gitlab client with the provided token.
		s.client, err = gitlab.NewClient(token)
		if err != nil {
			return fmt.Errorf("failed to create Gitlab client: %w", err)
		}

		if s.paginationMaxLimit == 0 {
			s.paginationMaxLimit = DefaultPaginationMaxLimit
		}

		// Set list opts for pagination.
		if s.listOptions == nil {
			s.listOptions = &gitlab.ListOptions{
				PerPage: s.paginationMaxLimit,
				Page:    DefaultPage,
				Sort:    Sort,
			}
		}

		return nil
	}
}

// WithPaginationMaxLimit sets the maximum number of items to request per page when paginating through results.
func WithPaginationMaxLimit(limit int) Options {
	return func(s *Service) error {
		s.paginationMaxLimit = limit
		return nil
	}
}
