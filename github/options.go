package github

import (
	"log/slog"
	"os"

	"github.com/cloudnative-zoo/go-commons/util"
	"github.com/gofri/go-github-ratelimit/github_ratelimit"
	"github.com/google/go-github/v67/github"
)

type Option func(*Service) error

func WithToken(token string) Option {
	return func(s *Service) error {
		if token == "" {
			token := util.GetEnv("GH_TOKEN", "GITHUB_TOKEN", "GITHUB_API_TOKEN", "GITHUB_OAUTH_TOKEN")
			if token == "" {
				slog.Error("Either token must be provided or any of GH_TOKEN, GITHUB_TOKEN, GITHUB_API_TOKEN, GITHUB_OAUTH_TOKEN environment variable must be set.")
				os.Exit(1)
			}
		}

		rateLimiter, err := github_ratelimit.NewRateLimitWaiterClient(nil)
		if err != nil {
			slog.Error("Failed to create github rate limiter client with error: %v", slog.Any("error", err))
			os.Exit(1)
		}
		s.client = github.NewClient(rateLimiter).WithAuthToken(token)
		return nil
	}
}
