package genai

import (
	"context"
	"errors"
	"fmt"

	"github.com/openai/openai-go"
)

type ChatCompletionResponse struct {
	Answer string                 `json:"answer"`
	Usage  openai.CompletionUsage `json:"usage"`
}

// ChatCompletion sends a message to the generative model.
func (s *Service) ChatCompletion(ctx context.Context, messages []openai.ChatCompletionMessageParamUnion) (ChatCompletionResponse, error) {
	response := ChatCompletionResponse{}
	params := openai.ChatCompletionNewParams{
		Messages: openai.F(messages),
		Model:    openai.F(s.config.Model),
		Seed:     openai.Int(1),
	}

	completion, err := s.client.Chat.Completions.New(ctx, params)

	if err != nil {
		return response, fmt.Errorf("failed to generate completion: %w", err)
	}

	if len(completion.Choices) == 0 {
		return response, errors.New("no completion choices returned")
	}

	response.Answer = completion.Choices[0].Message.Content
	response.Usage = completion.Usage
	return response, nil
}
