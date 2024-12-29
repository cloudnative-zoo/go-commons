package main

import (
	"context"
	"fmt"
	"log/slog"
	"os/exec"
	"strings"

	"github.com/cloudnative-zoo/go-commons/gemini"
	"github.com/google/generative-ai-go/genai"
)

func main() {
	ctx := context.Background()
	// Get git changes
	cmd := exec.Command("git", "diff", "--cached", "--name-status")
	changes, err := cmd.Output()
	if err != nil {
		slog.With("error", err).Error("failed to get git changes")
	}

	changesStr := strings.TrimSpace(string(changes))
	if changesStr == "" {
		slog.Info("no changes to commit")
		return
	}
	geminiClient, err := gemini.New(ctx, gemini.WithAPIKey("")) // export GEMINI_API_KEY=
	defer func(geminiClient *gemini.Service) {
		err := geminiClient.Close()
		if err != nil {
			slog.With("error", err).Error("failed to close gemini client")
		}
	}(geminiClient)
	if err != nil {
		slog.With("error", err).Error("failed to create gemini client")
	}
	resp, err := geminiClient.SendMessage(ctx, &gemini.SendMessageRequest{
		Model: "gemini-1.5-flash",
		Content: []*genai.Content{
			{
				Parts: []genai.Part{
					genai.Text("Provide a detailed commit message with a title and description. The title should be a concise summary (max 50 characters). The description should provide more context about the changes, explaining why the changes were made and their impact. Use bullet points if multiple changes are significant. If it's just some minor changes, use 'fix' instead of 'feat'. Do not include any explanation in your response, only return a commit message content."),
				},
				Role: "user",
			},
			{
				Parts: []genai.Part{
					genai.Text(fmt.Sprintf("Generate a commit message in conventional commit standard format based on the following file changes:\n\n%s\n\n- Commit message title must be a concise summary (max 100 characters)\n- If it's just some minor changes, use 'fix' instead of 'feat'\n- IMPORTANT: Do not include any explanation in your response, only return a commit message content", changesStr)),
				},
				Role: "user",
			},
		},
	})
	if err != nil {
		slog.With("error", err).Error("failed to send message to generative model")
	}
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				slog.With("part", part).Info("generated commit message")
			}
		}
	}

}
