// Package genai provides a client for using the Provide API parameters.
package genai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

// New initializes a new Service, applying any options and configuring the client.
func New(ctx context.Context, isAzure bool, opts ...Options) (*Service, error) {
	s := &Service{}

	// Apply each option to configure the service.
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	// Declare config ahead of the if/else.
	var config openai.ClientConfig

	if isAzure {
		// Initialize the Azure client.
		config = openai.DefaultAzureConfig(s.apiKey, s.baseURL)
		config.APIVersion = s.apiVersion

		// If you use a deployment name different from the model name, you can customize the AzureModelMapperFunc function
		config.AzureModelMapperFunc = func(model string) string {
			azureModelMapping := map[string]string{
				s.model: s.model,
			}
			return azureModelMapping[model]
		}
	} else {
		// Initialize the OpenAI client.
		config = openai.DefaultConfig(s.apiKey)
		// Optionally, set BaseURL if your service requires it.
		config.BaseURL = s.baseURL
	}

	client := openai.NewClientWithConfig(config)
	s.client = client

	return s, nil
}
