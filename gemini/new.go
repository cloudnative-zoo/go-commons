// Package gemini provides a client for using the Gemini API.
package gemini

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func New(ctx context.Context, opts ...Options) (*Service, error) {
	s := &Service{}

	// Apply each option to configure the service.
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(s.apiKey))
	if err != nil {
		log.Fatal(err)
	}
	s.client = client

	return s, nil
}
