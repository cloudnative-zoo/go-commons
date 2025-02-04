package genai

import (
	"github.com/sashabaranov/go-openai"
)

type Service struct {
	client     *openai.Client
	model      string
	apiKey     string
	baseURL    string
	apiVersion string
}
