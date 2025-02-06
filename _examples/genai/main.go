package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/sashabaranov/go-openai"

	"github.com/cloudnative-zoo/go-commons/genai"
	"github.com/cloudnative-zoo/go-commons/git"
)

type AIProvider string

const (
	Gemini      AIProvider = "gemini"
	DeepSeek    AIProvider = "deepseek"
	AzureOpenAI AIProvider = "azure-openai"
)

type ProviderConfig struct {
	APIKeyEnvVar string
	DefaultModel string
	BaseAPIURL   string
	APIVersion   string
}

var providerConfigs = map[AIProvider]ProviderConfig{
	Gemini: {
		APIKeyEnvVar: "GEMINI_API_KEY",
		DefaultModel: "gemini-2.0-flash-lite-preview-02-05",
		BaseAPIURL:   "https://generativelanguage.googleapis.com/v1beta/openai/",
	},
	DeepSeek: {
		APIKeyEnvVar: "DEEPSEEK_API_KEY",
		DefaultModel: "deepseek-chat",
		BaseAPIURL:   "https://api.deepseek.com/v1",
	},
	AzureOpenAI: {
		APIKeyEnvVar: "AZURE_OPENAI_API_KEY",
		DefaultModel: "o3-mini",
		BaseAPIURL:   "https://swedencentral.api.cognitive.microsoft.com/",
		APIVersion:   "2024-12-01-preview",
	},
}

type CommitGenerator struct {
	aiClient *genai.Service
}

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Initialize services
	gitSvc, err := setupGitService(ctx)
	if err != nil {
		logger.Error("Git service initialization failed", "error", err)
		os.Exit(1)
	}

	changes, err := gitSvc.Status()
	if err != nil {
		logger.Error("Failed to get repository status", "error", err)
		os.Exit(1)
	}

	if !hasChanges(changes) {
		logger.Info("No changes detected in repository")
		return
	}

	generator, err := NewCommitGenerator(ctx, AzureOpenAI)
	if err != nil {
		logger.Error("Failed to initialize commit generator", "error", err)
		os.Exit(1)
	}

	if err := generator.GenerateAndDisplayCommit(ctx, changes); err != nil {
		logger.Error("Commit generation failed", "error", err)
		os.Exit(1)
	}
}

func NewCommitGenerator(ctx context.Context, provider AIProvider) (*CommitGenerator, error) {
	config, ok := providerConfigs[provider]
	if !ok {
		return nil, fmt.Errorf("unsupported provider: %s. Supported providers: %v",
			provider, getSupportedProviders())
	}

	apiKey := os.Getenv(config.APIKeyEnvVar)
	if apiKey == "" {
		return nil, fmt.Errorf("missing required environment variable: %s", config.APIKeyEnvVar)
	}
	var isAzureOpenAI bool
	if provider == AzureOpenAI {
		isAzureOpenAI = true
	} else {
		isAzureOpenAI = false
	}
	aiClient, err := genai.New(
		ctx,
		isAzureOpenAI,
		genai.WithAPIKey(apiKey),
		genai.WithModel(config.DefaultModel),
		genai.WithBaseURL(config.BaseAPIURL),
		genai.WithAPIVersion(config.APIVersion),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create AI client: %w", err)
	}

	return &CommitGenerator{
		aiClient: aiClient,
	}, nil
}

func (cg *CommitGenerator) GenerateAndDisplayCommit(ctx context.Context, changes *git.StatusChanges) error {
	prompt := buildCommitPrompt(changes)
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}
	response, err := cg.aiClient.GenerateCompletion(ctx, messages)
	if err != nil {
		return fmt.Errorf("AI request failed: %w", err)
	}

	fmt.Printf("\nGenerated Commit:\n%s\n", formatResponse(response))
	return nil
}

func buildCommitPrompt(changes *git.StatusChanges) string {
	const promptTemplate = `Generate a Conventional Commits message and branch name:

Changes:
- Modified: %s
- Added:    %s
- Removed:  %s

Format:
"""
<type>[!]: <subject>

<body>

---
branch: <type>/<short-description>
"""

Requirements:
- Types: fix, feat, chore
- Subject: <=50 chars, imperative mood
- Body: Bullet points explaining changes
- Breaking changes: Append '!' and BREAKING CHANGE note
- Branch: 2-4 hyphenated words`

	return fmt.Sprintf(promptTemplate,
		strings.Join(changes.Modified, ", "),
		strings.Join(changes.Added, ", "),
		strings.Join(changes.Deleted, ", "),
	)
}

func setupGitService(ctx context.Context) (*git.Service, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current directory: %w", err)
	}

	svc, err := git.New(
		ctx,
		git.WithSSHKeyPath(path.Join(homeDir, ".ssh", "github_hassnatahmad"), ""),
		git.WithRepoPath(cwd),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create git service: %w", err)
	}
	return svc, nil
}

func hasChanges(changes *git.StatusChanges) bool {
	return len(changes.Added)+len(changes.Modified)+len(changes.Deleted) > 0
}

func formatResponse(response interface{}) string {
	str := fmt.Sprintf("%v", response)
	return strings.TrimSpace(str)
}

func getSupportedProviders() []string {
	providers := make([]string, 0, len(providerConfigs))
	for p := range providerConfigs {
		providers = append(providers, string(p))
	}
	return providers
}
