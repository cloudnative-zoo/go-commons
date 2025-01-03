package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/cloudnative-zoo/go-commons/genai/gemini"

	"github.com/cloudnative-zoo/go-commons/git"
)

// Constants
const (
	githubOrg  = "cloudnative-zoo"
	githubRepo = "go-commons"
)

func main() {
	ctx := context.Background()

	// Initialize services
	changes, err := initializeGitService(ctx)
	if err != nil {
		slog.With("error", err).Error("failed to initialize git service")
		return
	}

	if changes == nil || noChanges(changes) {
		slog.Info("repository is clean")
		return
	}

	geminiClient, err := initializeGeminiClient(ctx)
	if err != nil {
		slog.With("error", err).Error("failed to initialize gemini client")
		return
	}
	defer cleanupGeminiClient(geminiClient)
	// Generate and process commit message
	if err := generateCommitMessageWithBranch(ctx, geminiClient, changes); err != nil {
		slog.With("error", err).Error("failed to generate commit message")
	}
}

func initializeGitService(ctx context.Context) (*git.StatusChanges, error) {
	homeDir := os.Getenv("HOME")
	repoPath := path.Join(homeDir, "development", "github", githubOrg, githubRepo)

	// Initialize Git service
	gitSvc, err := git.New(
		ctx,
		git.WithSSHKeyPath(path.Join(homeDir, ".ssh", "github_hassnatahmad"), ""),
		git.WithRepoPath(repoPath),
		git.WithURL(fmt.Sprintf("git@github.com:%s/%s.git", githubOrg, githubRepo)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create git service: %w", err)
	}

	// Get git status
	changes, err := gitSvc.Status()
	if err != nil {
		return nil, fmt.Errorf("failed to get git status: %w", err)
	}

	return changes, nil
}

func noChanges(changes *git.StatusChanges) bool {
	return len(changes.Added) == 0 && len(changes.Modified) == 0 && len(changes.Deleted) == 0
}

func initializeGeminiClient(ctx context.Context) (*gemini.Service, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY is not set")
	}

	geminiClient, err := gemini.New(ctx, gemini.WithAPIKey(apiKey), gemini.WithModel("gemini-1.5-flash"))
	if err != nil {
		return nil, fmt.Errorf("failed to create gemini client: %w", err)
	}
	return geminiClient, nil
}

func cleanupGeminiClient(client *gemini.Service) {
	if client != nil {
		if err := client.Close(); err != nil {
			slog.With("error", err).Error("failed to close gemini client")
		}
	}
}

func generateCommitMessageWithBranch(ctx context.Context, geminiClient *gemini.Service, changes *git.StatusChanges) error {
	// Prepare the commit message and branch name prompt
	prompt := fmt.Sprintf(
		`Generate the following:
1. A commit message based on the following file changes:
Modified:
%s

Added:
%s

Deleted:
%s

The commit message must adhere to the conventional commit standard and Semantic Versioning (SemVer) principles:
- Use "fix" for bug fixes (patch).
- Use "feat" for new features (minor).
- Use "feat!:, fix!:, etc." for breaking changes (major).
- Ensure the title is concise (max 50 characters) and the body is detailed, explaining what and why.
- Ensure there are no placeholder values in the commit message.
- Ensure the commit message is in the imperative mood.
- Ensure that no "may", "might", "could", or "should" are used in the commit message.
- Ensure that the commit message has only bullet points, if necessary.

2. A suggested branch name for these changes:
- The branch name should use the format: [type]/[short-description].
- Use "feat", "fix", or "refactor" as the type, depending on the changes.
- The description should summarize the changes succinctly, using hyphens instead of spaces.
- Ensure the branch name is concise and descriptive.

Return only the formatted commit message and branch name.`,
		strings.Join(changes.Modified, "\n"),
		strings.Join(changes.Added, "\n"),
		strings.Join(changes.Deleted, "\n"),
	)

	// Send the prompt to Gemini
	/*	resp, err := geminiClient.SendMessage(ctx, &gemini.SendMessageRequest{
		Model: "gemini-1.5-flash",
		Content: []*genai.Content{
			{
				Parts: []genai.Part{
					genai.Text(prompt),
				},
				Role: "user",
			},
		},
	})*/
	resp, err := geminiClient.GenerateCompletion(ctx, prompt)
	if err != nil {
		return fmt.Errorf("failed to send message to Gemini: %w", err)
	}

	formattedMessage := strings.TrimSpace(fmt.Sprintf("%v", resp))
	slog.Info("\nGenerated Output:\n" + formattedMessage)

	return nil
}
