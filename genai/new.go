// Package genai provides a client for using the Provide API parameters.
package genai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func New(ctx context.Context, opts ...Options) (*Service, error) {
	s := &Service{}

	// Apply each option to configure the service.
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}
	config := openai.DefaultConfig(s.apiKey)
	config.BaseURL = s.baseURL

	client := openai.NewClientWithConfig(config)
	s.client = client

	return s, nil
}
