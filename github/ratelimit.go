package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v69/github"
)

// CheckRateLimit retrieves the current rate limits for the authenticated user or application.
// Returns detailed rate limit information for core, search, and other GitHub API categories.
func (s *Service) CheckRateLimit(ctx context.Context) (*github.RateLimits, error) {
	rateLimits, _, err := s.client.RateLimit.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve rate limits: %w", err)
	}
	return rateLimits, nil
}
