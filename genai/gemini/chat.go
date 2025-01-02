package gemini

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/generative-ai-go/genai"
)

// GetCompletion sends a message to the generative model.
func (s *Service) GetCompletion(ctx context.Context, prompt string) (string, error) {
	model := s.client.GenerativeModel(s.model)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}
	if len(resp.Candidates) == 0 {
		if resp.PromptFeedback.BlockReason == genai.BlockReasonSafety {
			for _, r := range resp.PromptFeedback.SafetyRatings {
				if !r.Blocked {
					continue
				}
				return "", fmt.Errorf("completion blocked due to %v with probability %v", r.Category.String(), r.Probability.String())
			}
		}
		return "", errors.New("no completion returned; unknown reason")
	}
	got := resp.Candidates[0]
	var output string
	for _, part := range got.Content.Parts {
		switch o := part.(type) {
		case genai.Text:
			output += string(o)
			output += "\n"
		default:
			return "", fmt.Errorf("unexpected part type %T", o)
		}
	}

	if got.CitationMetadata != nil && len(got.CitationMetadata.CitationSources) > 0 {
		output += "Citations:\n"
		for _, source := range got.CitationMetadata.CitationSources {
			output += fmt.Sprintf("* %s, %s\n", *source.URI, source.License)
		}
	}
	return output, nil
}
