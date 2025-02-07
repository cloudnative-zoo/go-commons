package genai

import (
	"errors"
	"os"
)

// Option defines a function signature for configuring a Service instance.
type Option func(*Config) error

// WithAPIKey configures the service with the provided API key.
func WithAPIKey(apiKey string) Option {
	return func(c *Config) error {
		if apiKey == "" {
			apiKey = os.Getenv("GENAI_API_KEY")
			if apiKey == "" {
				return errors.New("API key cannot be empty")
			}
		}
		c.APIKey = apiKey
		return nil
	}
}

// WithProvider sets the AI provider.
func WithProvider(provider Provider) Option {
	return func(c *Config) error {
		c.Provider = provider
		return nil
	}
}

// WithModel overrides the default model.
func WithModel(model string) Option {
	return func(c *Config) error {
		if model == "" {
			return errors.New("model cannot be empty")
		}
		c.Model = model
		return nil
	}
}

// WithBaseURL overrides the default base URL.
func WithBaseURL(baseURL string) Option {
	return func(c *Config) error {
		if baseURL == "" {
			return errors.New("baseURL cannot be empty")
		}
		c.BaseURL = baseURL
		return nil
	}
}

// WithAPIVersion sets the API version (required for Azure).
func WithAPIVersion(apiVersion string) Option {
	return func(c *Config) error {
		if c.Provider == ProviderAzureOpenAI && apiVersion == "" {
			return errors.New("apiVersion is required for Azure OpenAI")
		}
		c.APIVersion = apiVersion
		return nil
	}
}
