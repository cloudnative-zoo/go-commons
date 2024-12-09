package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v67/github"
)

func (s *Service) CheckRateLimit(ctx context.Context) (*github.RateLimits, error) {
	rateLimits, _, err := s.client.RateLimit.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve rate limits with error: %w", err)
	}
	return rateLimits, nil
}
