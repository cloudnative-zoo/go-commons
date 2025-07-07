package github

import (
	"errors"

	"github.com/cloudnative-zoo/go-commons/utilities"
	"github.com/gofri/go-github-ratelimit/v2/github_ratelimit"
	"github.com/google/go-github/v73/github"
)

// Options defines a function signature for configuring a GitHub Service instance.
type Options func(*Service) error

// WithToken sets up authentication for the GitHub client using a personal access token.
// The token can be provided directly or sourced from environment variables.
func WithToken(token string) Options {
	return func(s *Service) error {
		if token == "" {
			// Fetch token from environment variables if not provided.
			token = utilities.GetEnv("GH_TOKEN", "GITHUB_TOKEN", "GITHUB_API_TOKEN", "GITHUB_OAUTH_TOKEN")
			if token == "" {
				return errors.New("GitHub token is required: provide a token or set the GH_TOKEN, GITHUB_TOKEN, GITHUB_API_TOKEN, or GITHUB_OAUTH_TOKEN environment variable")
			}
		}

		// Create a rate limiter-enabled GitHub client.
		rateLimiter := github_ratelimit.NewClient(nil)

		// Initialize the GitHub client with the provided token.
		s.client = github.NewClient(rateLimiter).WithAuthToken(token)
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
