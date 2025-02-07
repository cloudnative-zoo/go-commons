package genai

import (
	"context"
	"errors"
	"fmt"

	"github.com/openai/openai-go"
)

// GenerateCompletion sends a message to the generative model.
func (s *Service) GenerateCompletion(ctx context.Context, messages []openai.ChatCompletionMessageParamUnion) (string, error) {
	params := openai.ChatCompletionNewParams{
		Messages: openai.F(messages),
		Model:    openai.F(s.config.Model),
		Seed:     openai.Int(1),
	}

	completion, err := s.client.Chat.Completions.New(ctx, params)

	if err != nil {
		return "", fmt.Errorf("failed to generate completion: %w", err)
	}

	if len(completion.Choices) == 0 {
		return "", errors.New("no completion choices returned")
	}

	return completion.Choices[0].Message.Content, nil
}
