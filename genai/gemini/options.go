package gemini

import (
	"errors"
	"os"
)

// Options defines a function signature for configuring a GitHub Service instance.
type Options func(*Service) error

// WithAPIKey configures the Gemini client with the provided apiKey.
func WithAPIKey(apiKey string) Options {
	return func(s *Service) error {
		if apiKey == "" {
			apiKey = os.Getenv("GEMINI_API_KEY")
			if apiKey == "" {
				return errors.New("apiKey cannot be empty. Set GEMINI_API_KEY environment variable or provide a value")
			}
		}
		s.apiKey = apiKey // pragma: allowlist secret
		return nil
	}
}
