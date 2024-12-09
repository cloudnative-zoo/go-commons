package github

import (
	"log/slog"
	"os"

	"github.com/cloudnative-zoo/go-commons/util"
	"github.com/gofri/go-github-ratelimit/github_ratelimit"
	"github.com/google/go-github/v67/github"
)

type Option func(*Service) error

// WithToken sets up authentication for the GitHub client using a personal access token.
// The token can be provided directly or sourced from environment variables.
func WithToken(token string) Option {
	return func(s *Service) error {
		if token == "" {
			// Fetch token from environment variables if not provided.
			token = util.GetEnv("GH_TOKEN", "GITHUB_TOKEN", "GITHUB_API_TOKEN", "GITHUB_OAUTH_TOKEN")
			if token == "" {
				slog.Error("A valid token must be provided directly or via environment variables (GH_TOKEN, GITHUB_TOKEN, GITHUB_API_TOKEN, GITHUB_OAUTH_TOKEN).")
				os.Exit(1)
			}
		}

		// Create a rate limiter-enabled GitHub client.
		rateLimiter, err := github_ratelimit.NewRateLimitWaiterClient(nil)
		if err != nil {
			slog.With("error", err).Error("Failed to create GitHub rate limiter client")
			os.Exit(1)
		}

		// Initialize the GitHub client with the provided token.
		s.client = github.NewClient(rateLimiter).WithAuthToken(token)
		return nil
	}
}
