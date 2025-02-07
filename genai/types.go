package genai

import (
	"github.com/openai/openai-go"
)

// Provider represents supported AI providers.
type Provider string

const (
	ProviderOpenAI      Provider = "openai"
	ProviderAzureOpenAI Provider = "azure-openai"
)

// Service handles AI completions and configurations.
type Service struct {
	client *openai.Client
	config Config
}

// Config holds service configuration.
type Config struct {
	Provider   Provider
	Model      string
	APIKey     string
	BaseURL    string
	APIVersion string
}
