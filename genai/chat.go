package genai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

// GenerateCompletion sends a message to the generative model.
func (s *Service) GenerateCompletion(ctx context.Context, messages []openai.ChatCompletionMessage) (string, error) {
	resp, err := s.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    s.model,
			Messages: messages,
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate completion: %w", err)
	}
	return resp.Choices[0].Message.Content, err
}
