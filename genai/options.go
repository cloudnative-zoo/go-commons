package genai

import (
	"errors"
	"os"
)

// Options defines a function signature for configuring a GitHub Service instance.
type Options func(*Service) error

// WithAPIKey configures the OpenAI client with the provided apiKey.
func WithAPIKey(apiKey string) Options {
	return func(s *Service) error {
		if apiKey == "" {
			apiKey = os.Getenv("GENAI_API_KEY")
			if apiKey == "" {
				return errors.New("apiKey cannot be empty. Set GENAI_API_KEY environment variable or provide a value")
			}
		}
		s.apiKey = apiKey // pragma: allowlist secret
		return nil
	}
}

// WithModel configures the OpenAI client with the provided model.
func WithModel(model string) Options {
	return func(s *Service) error {
		if model == "" {
			return errors.New("model cannot be empty")
		}
		s.model = model
		return nil
	}
}

// WithBaseURL configures the OpenAI client with the provided baseURL.
func WithBaseURL(baseURL string) Options {
	return func(s *Service) error {
		if baseURL == "" {
			return errors.New("baseURL cannot be empty")
		}
		s.baseURL = baseURL
		return nil
	}
}

// WithAPIVersion configures the Azure client with the provided apiVersion.
func WithAPIVersion(apiVersion string) Options {
	return func(s *Service) error {
		s.apiVersion = apiVersion
		return nil
	}
}
