package gemini

import (
	"context"

	"github.com/google/generative-ai-go/genai"
)

type SendMessageRequest struct {
	Model   string
	Content []*genai.Content
}

// SendMessage sends a message to the generative model.
func (s *Service) SendMessage(ctx context.Context, req *SendMessageRequest) (*genai.GenerateContentResponse, error) {
	model := s.client.GenerativeModel(req.Model)
	cs := model.StartChat()
	cs.History = req.Content
	resp, err := cs.SendMessage(ctx, genai.Text(""))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
