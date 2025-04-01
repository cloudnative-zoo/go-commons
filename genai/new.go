// Package genai provides a client for using the Provide API parameters.
package genai

import (
	"fmt"

	"github.com/openai/openai-go/azure"
	"github.com/openai/openai-go/option"

	"github.com/openai/openai-go"
)

// New initializes a new Service with the provided options.
func New(opts ...Option) (*Service, error) {
	config := &Config{
		Provider: ProviderOpenAI, // Default provider
	}

	// Apply user options
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}
	var client openai.Client
	// Create client
	if config.Provider == ProviderAzureOpenAI {
		client = openai.NewClient(
			azure.WithEndpoint(config.BaseURL, config.APIVersion),
			azure.WithAPIKey(config.APIKey),
		)
	} else {
		client = openai.NewClient(
			option.WithBaseURL(config.BaseURL),
			option.WithAPIKey(config.APIKey),
		)
	}

	return &Service{
		client: &client,
		config: *config,
	}, nil
}
