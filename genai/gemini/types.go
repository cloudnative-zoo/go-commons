package gemini

import (
	"errors"

	"github.com/google/generative-ai-go/genai"
)

type Service struct {
	client *genai.Client
	model  string
	apiKey string
}

func (s *Service) Close() error {
	err := s.client.Close()
	if err != nil {
		return errors.New("failed to close client")
	}
	return nil
}
